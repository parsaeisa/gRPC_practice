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
### Implement Server 
At first you should define a server struct which has a UnimplementedServer from generated files .
Then you start the server listening. 
for each of services that you define in your .proto files you should define methods and assign them to Server struct . 


### Implement Client 
Implementing client is easy . you just define one method to make connection to server , then all you have to do is calling methods that you defined for the server . 

## Streaming with gRPC
In the previous section , we call one of the server functions and server responds to us immediately . 
But how about when we (as a client ) want the server to notify us when something special happens ? 
here we need gRPC streaming . 

like previous part you should define messages and services , with little difference :
```proto

```

## Client 
in gRPC streaming , implementing client has more challenges than implementing server . 
In Client , after connecting to server and define a stream in the way below :
```go 
    // dial to server 
    conn , err := grpc.Dial(address , grpc.WithInsecure() , grpc.WithBlock())
    if err != nil {
      log.Fatalf("could not connect %v" , err)
    }	
    defer conn.Close()
    
    // make client
    client := pb.NewGrpcClient(conn)    
		ctx := context.Background()    
		stream, err := client.NatsSubscribe(ctx)
		if err != nil {
			log.WithError(err).Fatal("error in sending request")
		}
```
We have two go-routins wich operate :
+ sending requests to server
+ waiting for responses from server

### Server 
but in server we only have one go-routin . 
> **NOTE:** This example is so simple . based on you project you can define many go-routins . 

> **NOTE:** you can read about other kinds of gRPC streaming in [here](https://grpc.io/docs/languages/go/basics/#defining-the-service)
