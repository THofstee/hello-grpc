package main

import (
    "context"
    "flag"
    "log"
    "time"
    "google.golang.org/grpc"
    pb "hello_grpc/interface"
)

var (
    serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

func main() {
    flag.Parse()
    var opts []grpc.DialOption
    opts = append(opts, grpc.WithInsecure())

    conn, err := grpc.Dial(*serverAddr, opts...)
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewHelloClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    _, err = client.Log(ctx, &pb.String{Msg: "Hello, world!"})
    if err != nil {
        log.Fatalf("failed to log: %v", err)
    }
}
