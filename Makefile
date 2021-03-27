all: build

backend:
	go build -o out/liwasc-backend/liwasc-backend cmd/liwasc-backend/main.go

frontend:
	rm -f web/app.wasm
	GOOS=js GOARCH=wasm go build -o web/app.wasm cmd/liwasc-frontend/main.go
	go build -o /tmp/liwasc-frontend-build cmd/liwasc-frontend/build.go
	rm -rf out/liwasc-frontend
	/tmp/liwasc-frontend-build -build
	cp -r web/* out/liwasc-frontend/web

build: backend frontend

release-backend:
	CGO_ENABLED=1 go build -ldflags="-extldflags=-static" -tags sqlite_omit_load_extension -o out/release/liwasc-backend/liwasc-backend.linux-$$(uname -m) cmd/liwasc-backend/main.go

release-frontend: frontend
	rm -rf out/release/liwasc-frontend
	mkdir -p out/release/liwasc-frontend
	cd out/liwasc-frontend && tar -czvf ../release/liwasc-frontend/liwasc-frontend.tar.gz .

release-frontend-github-pages: frontend
	rm -rf out/release/liwasc-frontend-github-pages
	mkdir -p out/release/liwasc-frontend-github-pages
	/tmp/liwasc-frontend-build -build -path liwasc -out out/release/liwasc-frontend-github-pages
	cp -r web/* out/release/liwasc-frontend-github-pages/web

release: release-backend release-frontend release-frontend-github-pages

dev:
	while [ -z "$$BACKEND_PID" ] || [ -n "$$(inotifywait -q -r -e modify pkg cmd web/*.css)" ]; do\
		$(MAKE);\
		sudo pkill -9 -P $$BACKEND_PID 2>/dev/null 1>&2;\
		kill -9 $$FRONTEND_PID 2>/dev/null 1>&2;\
		sudo out/liwasc-backend/liwasc-backend -oidcIssuer=${OIDCISSUER} -oidcClientID=${OIDCCLIENTID} -deviceName=${DEVICENAME} & export BACKEND_PID="$$!";\
		/tmp/liwasc-frontend-build -serve & export FRONTEND_PID="$$!";\
	done

clean:
	rm -rf out
	rm -rf pkg/proto/generated
	rm -rf pkg/databases/generated

depend:
	# Initialize working directories
	sudo mkdir -p /etc/liwasc /var/lib/liwasc
	sudo chown -R $$USER /etc/liwasc /var/lib/liwasc
	# Download external databases
	curl -L -o /etc/liwasc/oui-database.sqlite https://mac2vendor.com/download/oui-database.sqlite
	curl -L -o /etc/liwasc/service-names-port-numbers.csv https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.csv
	curl -L -o /etc/liwasc/ports2packets.csv https://github.com/pojntfx/ports2packets/releases/download/weekly-csv/ports2packets.csv
	# Get CLIs
	GO111MODULE=on go get github.com/volatiletech/sqlboiler/v4@latest
	GO111MODULE=on go get github.com/volatiletech/sqlboiler-sqlite3@latest
	GO111MODULE=on go get github.com/golang/protobuf/protoc-gen-go@latest
	GO111MODULE=on go get github.com/rubenv/sql-migrate/@latest
	GO111MODULE=on go get github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	# Migrate databases manually
	sql-migrate down -env="production" -config pkg/databases/node_and_port_scan.yaml
	sql-migrate down -env="production" -config pkg/databases/node_wake.yaml
	sql-migrate up -env="production" -config pkg/databases/node_and_port_scan.yaml
	sql-migrate up -env="production" -config pkg/databases/node_wake.yaml
	# Generate database and API bindings
	go generate ./...
	# Build
	go get ./...
	