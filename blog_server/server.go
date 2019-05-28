package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"../blogpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Blog Service Started!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting server . . .")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve : %v", err)
		}
	}()

	// Wait for Ctrl-C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	// Block until a signal is received
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Stopping the listener")
	lis.Close()
}
