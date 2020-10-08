build:
	@go build -o liwasc main.go

run: build
	@./liwasc -oidcIssuer=${OIDCISSUER} -oidcClientID=${OIDCCLIENTID} -deviceName=${DEVICENAME}
