
build-linux:
	CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o trains_linux main.go

build-mac:
	CGO_ENABLED=0 GOOS=darwain go build -v -a -installsuffix cgo -o trains_mac main.go