package proto

//go:generate sh -c "mkdir -p generated && protoc --go_out=paths=source_relative,plugins=grpc:generated -I=../../api ../../api/*.proto"
