build:
	@GOARCH=wasm GOOS=js go build -o web/app.wasm cmd/liwasc-frontend-web-app/main.go
	@go build -o liwasc-frontend-web-server cmd/liwasc-frontend-web-server/main.go

build-static: build
	@rm -rf out
	@./liwasc-frontend-web-server -backendURL=${BACKEND_URL} -oidcIssuer=${OIDCISSUER} -oidcClientID=${OIDCCLIENTID} -oidcRedirectURL=${OIDCREDIRECTURL} -buildStatic
	@cp -r web/* out/web/

run: build
	@./liwasc-frontend-web-server -backendURL=${BACKEND_URL} -oidcIssuer=${OIDCISSUER} -oidcClientID=${OIDCCLIENTID} -oidcRedirectURL=${OIDCREDIRECTURL}

clean:
	@rm -rf ./pkg/proto/generated