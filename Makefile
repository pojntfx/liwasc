build:
	@go build -o liwasc main.go

run: build
	@sudo ./liwasc -oidcIssuer=${OIDCISSUER} -oidcClientID=${OIDCCLIENTID} -deviceName=${DEVICENAME}

clean:
	@rm -rf ./pkg/sql/generated
	@rm -rf ./pkg/proto/generated

deps:
	@go install github.com/volatiletech/sqlboiler/v4@latest
	@go install github.com/volatiletech/sqlboiler-sqlite3@latest
	@go install github.com/golang/protobuf/protoc-gen-go@latest

tools:
	@go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

setup-databases:
	@sudo mkdir -p /etc/liwasc
	@sudo curl -L -o /etc/liwasc/oui-database.sqlite https://mac2vendor.com/download/oui-database.sqlite
	@sudo curl -L -o /etc/liwasc/service-names-port-numbers.csv https://www.iana.org/assignments/service-names-port-numbers/service-names-port-numbers.csv
	@sudo curl -L -o /etc/liwasc/ports2packets.csv https://github.com/pojntfx/ports2packets/releases/download/weekly-csv/ports2packets.csv
	@sudo mkdir -p /var/liwasc
	@sudo sqlite3 -batch /var/liwasc/node_and_port_scan.sqlite ".read ./pkg/sql/node_and_port_scan.sql"
	@sudo sqlite3 -batch /var/liwasc/node_wake.sqlite ".read ./pkg/sql/node_wake.sql"