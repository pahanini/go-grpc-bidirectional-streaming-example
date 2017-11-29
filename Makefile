
all: client server

dep:
	@echo "Install dependencies"
	cd src && glide install

protoc:
	@echo "Generating Go files"
	cd src/proto && protoc --go_out=plugins=grpc:. *.proto

server: protoc
	@echo "Building server"
	go build -i -o server server

client: protoc
	@echo "Building client"
	go build -i -o client client

.PHONY: client server protoc dep
