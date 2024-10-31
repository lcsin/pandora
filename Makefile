run:
	@go build -o pandora.exe
	@./pandora.exe
windows:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o pandora.exe main.go 
mac:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o pandora main.go 
linuxamd:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pandora main.go 
linuxarm:
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o pandora main.go 