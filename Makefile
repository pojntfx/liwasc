all: backend frontend

backend:
	go build -o out/liwasc-backend/liwasc-backend cmd/liwasc-backend/main.go

frontend:
	GOOS=js GOARCH=wasm go build -o web/app.wasm cmd/liwasc-frontend/main.go
	go build -o /tmp/liwasc-frontend-build cmd/liwasc-frontend/build.go
	rm -rf out/liwasc-frontend
	/tmp/liwasc-frontend-build -build
	cp -r web/* out/liwasc-frontend/web

dev:
	while [ -z "$$BACKEND_PID" ] || [ -n "$$(inotifywait -q -r -e modify pkg cmd)" ]; do\
		$(MAKE);\
		sudo pkill -9 -P $$BACKEND_PID 2>/dev/null 1>&2;\
		kill -9 $$FRONTEND_PID 2>/dev/null 1>&2;\
		sudo out/liwasc-backend/liwasc-backend -oidcIssuer=${OIDCISSUER} -oidcClientID=${OIDCCLIENTID} -deviceName=${DEVICENAME} & export BACKEND_PID="$$!";\
		/tmp/liwasc-frontend-build -serve & export FRONTEND_PID="$$!";\
	done

clean:
	rm -rf out
	rm -rf pkg/proto/generated
	rm -rf pkg/sql/generated

depend:
	sudo mkdir -p /etc/liwasc /var/lib/liwasc
	sudo chown -R $$USER /etc/liwasc /var/lib/liwasc
	curl -L -o /etc/liwasc/oui-database.sqlite https://mac2vendor.com/download/oui-database.sqlite
	curl -L -o /etc/liwasc/service-names-port-numbers.csv https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.csv
	curl -L -o /etc/liwasc/ports2packets.csv https://github.com/pojntfx/ports2packets/releases/download/weekly-csv/ports2packets.csv
	sqlite3 -batch /var/lib/liwasc/node_and_port_scan.sqlite ".read pkg/sql/node_and_port_scan.sql"
	sqlite3 -batch /var/lib/liwasc/node_wake.sqlite ".read pkg/sql/node_wake.sql"
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler-sqlite3@latest
	go install github.com/golang/protobuf/protoc-gen-go@latest
	go generate ./...
	go get ./...
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	