all: build

backend:
	go build -o out/liwasc cmd/liwasc-backend/main.go

wasm:
	GOARCH=wasm GOOS=js go build -o web/app.wasm cmd/liwasc-frontend/main.go

site: wasm
	rm -rf out
	go run cmd/liwasc-frontend-builder/main.go -build
	cp -r web/* out/web

build: wasm site backend

run-backend:
	sudo out/liwasc -oidcIssuer=${OIDCISSUER} -oidcClientID=${OIDCCLIENTID} -deviceName=${DEVICENAME}

serve: wasm
	go run cmd/liwasc-frontend-builder/main.go -serve

generate:
	go generate ./...

clean:
	rm -rf out
	rm -rf pkg/proto/generated
	rm -rf pkg/sql/generated

deps:
	go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
	go install github.com/volatiletech/sqlboiler/v4@latest
	go install github.com/volatiletech/sqlboiler-sqlite3@latest
	go install github.com/golang/protobuf/protoc-gen-go@latest
	sudo mkdir -p /etc/liwasc
	sudo curl -L -o /etc/liwasc/oui-database.sqlite https://mac2vendor.com/download/oui-database.sqlite
	sudo curl -L -o /etc/liwasc/service-names-port-numbers.csv https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.csv
	sudo curl -L -o /etc/liwasc/ports2packets.csv https://github.com/pojntfx/ports2packets/releases/download/weekly-csv/ports2packets.csv
	sudo mkdir -p /var/liwasc
	sudo sqlite3 -batch /var/liwasc/node_and_port_scan.sqlite ".read ./pkg/sql/node_and_port_scan.sql"
	sudo sqlite3 -batch /var/liwasc/node_wake.sqlite ".read ./pkg/sql/node_wake.sql"
