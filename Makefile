all: build

wasm:
	GOARCH=wasm GOOS=js go build -o web/app.wasm cmd/liwasc-frontend/main.go

site: wasm
	rm -rf out
	go run cmd/liwasc-frontend-builder/main.go -build
	cp -r web/* out/web

build: wasm site

serve: wasm
	go run cmd/liwasc-frontend-builder/main.go -serve

generate:
	go generate ./...

clean:
	rm -rf out
	rm -rf pkg/proto/generated
