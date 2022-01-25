# gRPC_practice

## Why GRPC ? 

## getting started with grpc :
### Install dependencies 
through [here](https://grpc.io/docs/languages/go/quickstart/) you can see GRPC document for golang and other languages .
before getting started make sure you have installed protobuf-compiler . 


- Add Protoc to your PATH 
```shell script
export PATH="$PATH:$(go env GOPATH)/bin"
```

> **NOTE:** remember export doesn't add variable to PATH in another terminal . So you should use the command in the same terminal which you exported the variable . 

### Define protofiles
after adding these dependencies , you describe your server  with .proto files . here you can see a simple example : 
with "message" keyword you can describe messages between your server and client : 
```proto
message HelloRequest {
  string name = 1;
  string lastname = 2;
  uint64 age = 3 ;
}
```
and with "service" keyword you can describe the behavior of your Server :

``` proto
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```
### Generate codes

use this command to generate .pb.go files : 
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
```
