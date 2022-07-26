to generate the proto files from the first implementation dir:

```azure 
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    implementation/implementation.proto
```

then to run the server from the same directory type:

```azure
go run server/main.go 
```

and to run the client again from the same directory but different terminal:

```azure
    go run client/main.go -name="GOPATH"
```

adding the env variable you want in the name argument.
