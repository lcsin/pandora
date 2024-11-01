run:
	@go build -o pandora.exe
	@./pandora.exe
windows:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o pandora.exe
mac:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o pandora
linuxamd:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pandora
linuxarm:
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o pandora