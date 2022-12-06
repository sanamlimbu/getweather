# gRPC tutorial

Simple application that uses gRPC and protocol buffer. Server fetches weather information of given location from [Open-Meteo](https://open-meteo.com/en) and returns to the client.

### _Dev dependencies_

- [Go](https://go.dev/dl/)
- [Protocol buffer compiler](https://github.com/protocolbuffers/protobuf/releases/tag/v21.10)
- [Go plugins](https://grpc.io/docs/languages/go/quickstart/) for the protocol compiler

1.  Install the protocol compiler plugins for Go using the following commands:
    ```
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
    ```
2.  Update your PATH so that the protoc compiler can find the plugins:

    ```
    $ export PATH="$PATH:$(go env GOPATH)/bin"
    ```

    Note: Mac user don't follow step 2 and use following command:

    ```
    $ brew install protoc-gen-go
    $ brew install protoc-gen-go-grpc
    ```
