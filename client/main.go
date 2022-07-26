package main

import (
	"context"
	"flag"
	"log"
	"time"

	ser "github.com/gkyriazis-ionos/SimpleApplication/implementation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	//pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//c := pb.NewGreeterClient(conn)
	c := ser.NewGetEnvVarClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	r, err := c.GetEnvVar(ctx, &ser.WhichEnvVar{WhichEnvVar: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("The Env Var is: %s", r.GetEnvVar())
}
