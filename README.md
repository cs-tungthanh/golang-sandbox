# golang-sandbox
SandBox to test everything, infra code


# Installation

## Install protobuf

```
<!-- Install protobuf in your OS -->
brew install protobuf
<!-- We can check if it's working or not by running the protoc command -->

<!-- Install proto plugin for golang -->
<!-- https://grpc.io/docs/languages/go/quickstart/ -->
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```