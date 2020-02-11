# Golang gRPC bidirectional streaming example

- client sends random numbers to server
- server receives number and sends it back if the number greater than all previous numbers
- both client and serever handle context errors (try to close client during send)

## Requirements

- go 1.12
- protobuf installed
- go support for protobuf installed

## Installation

### MacOS

```bash
brew install go
brew install protobuf
go get -u github.com/golang/protobuf/protoc-gen-go
```

Make sure ```protoc-gen-go``` added in PATH

### Linux

TBD

## Complie

```bash
make all
```

It should create two binaries `server` and `client`

## Use

Start server `./server` and in other terminal start `./client`

Client output example:

```bash
./client
2017/12/01 14:16:54 0 sent
2017/12/01 14:16:54 1 sent
2017/12/01 14:16:54 new max 1 received
2017/12/01 14:16:55 2 sent
2017/12/01 14:16:55 new max 2 received
2017/12/01 14:16:55 0 sent
2017/12/01 14:16:55 0 sent
2017/12/01 14:16:55 4 sent
2017/12/01 14:16:55 new max 4 received
2017/12/01 14:16:55 0 sent
2017/12/01 14:16:56 6 sent
2017/12/01 14:16:56 new max 6 received
2017/12/01 14:16:56 3 sent
2017/12/01 14:16:56 2 sent
2017/12/01 14:16:56 finished with max=6
```

Server output:

```bash
./server
2017/12/01 14:16:54 start new server
2017/12/01 14:16:54 send new max=1
2017/12/01 14:16:55 send new max=2
2017/12/01 14:16:55 send new max=4
2017/12/01 14:16:56 send new max=6
2017/12/01 14:16:56 exit
````
