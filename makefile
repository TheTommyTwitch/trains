all:

run-linux:
	./trains_linux -stations=./stations.dat -trains=./trains.dat

build-linux:
	CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o trains_linux main.go

run-mac:
	./trains_mac -stations=./stations.dat -trains=./trains.dat

build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v -a -installsuffix cgo -o trains_mac main.go