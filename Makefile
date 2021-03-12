all: build

backend:
	go build -o out/backend/liwasc cmd/liwasc-backend/main.go

wasm:
	GOARCH=wasm GOOS=js go build -o web/app.wasm cmd/liwasc-frontend/main.go

frontend-builder: wasm
	go build -o out/frontend-builder/liwasc-frontend-builder cmd/liwasc-frontend-builder/main.go

frontend-static: frontend-builder wasm
	rm -rf out/frontend-static
	out/frontend-builder/liwasc-frontend-builder -build
	cp -r web/* out/frontend-static

build: backend frontend-builder frontend-static

dev:
	while inotifywait -r -e modify pkg cmd; do\
		# Build\
		$(MAKE) -j2 backend frontend-builder;\
		# Kill\
		sudo pkill -9 -P $$BACKEND_PID;\
		sudo pkill -9 -P $$FRONTEND_PID;\
		# Run\
		sudo out/backend/liwasc -oidcIssuer=${OIDCISSUER} -oidcClientID=${OIDCCLIENTID} -deviceName=${DEVICENAME} & export BACKEND_PID="$$!";\
		sudo out/frontend-builder/liwasc-frontend-builder -serve & export FRONTEND_PID="$$!";\
	done

generate:
	go generate ./...

clean:
	rm -rf out
	rm -rf pkg/proto/generated
	rm -rf pkg/sql/generated

depend:
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler-sqlite3@latest
	go install github.com/golang/protobuf/protoc-gen-go@latest
	sudo mkdir -p /etc/liwasc
	sudo curl -L -o /etc/liwasc/oui-database.sqlite https://mac2vendor.com/download/oui-database.sqlite
	sudo curl -L -o /etc/liwasc/service-names-port-numbers.csv https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.csv
	sudo curl -L -o /etc/liwasc/ports2packets.csv https://github.com/pojntfx/ports2packets/releases/download/weekly-csv/ports2packets.csv
	sudo mkdir -p /var/liwasc
	sudo sqlite3 -batch /var/liwasc/node_and_port_scan.sqlite ".read pkg/sql/node_and_port_scan.sql"
	sudo sqlite3 -batch /var/liwasc/node_wake.sqlite ".read pkg/sql/node_wake.sql"
