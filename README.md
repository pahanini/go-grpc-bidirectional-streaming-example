# gRPC bidirectional streaming example

In this example client sends random numbers to server and server sends back to client received number 
if this number bigger than all received before.

## Requirements

- go 1.9
- glide installed 
- protobuf installed
- go support for protobuf installed

## Installation

### MacOS

```
brew install go
brew install glide
brew install protobuf
go get -u github.com/golang/protobuf/protoc-gen-go 
```

Make sure ```protoc-gen-go``` added in PATH

### Linux 

TBD

## Complie

```
make dep
make all
```

It should create two binaries `server` and `client`

## Use

Start server `./server` and in other terminal start `./client`
