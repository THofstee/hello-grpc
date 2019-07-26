package main

import (
    "context"
    "flag"
    "fmt"
    "log"
    "net"
    "google.golang.org/grpc"
   pb "hello_grpc/interface"
)

var (
    port = flag.Int("port", 10000, "The server port")
)

type helloServer struct {}

func (s* helloServer) Log(ctx context.Context, string *pb.String) (*pb.Empty, error) {
    fmt.Println(string.Msg)
    return &pb.Empty{}, nil
}

func main() {
    flag.Parse()
    lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    pb.RegisterHelloServer(grpcServer, &helloServer{})
    grpcServer.Serve(lis)
}
