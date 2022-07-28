package main

import (
	"context"
	"flag"
	"fmt"
	ser "github.com/gkyriazis-ionos/SimpleApplication/implementation"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	ser.UnimplementedGetEnvVarServer
}

func (s *server) GetEnvVar(ctx context.Context, envVar *ser.WhichEnvVar) (*ser.EnvVar, error) {
	log.Printf("Received: %s", envVar.WhichEnvVar)
	return &ser.EnvVar{EnvVar: os.Getenv(envVar.WhichEnvVar)}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ser.RegisterGetEnvVarServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	log.Printf("Backend 1")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
