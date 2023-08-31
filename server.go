package main

import (
	"log"
	"net"

	"github.com/iuriimr/grpcecho/echo"
	grpc "google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("Listen TCP error: ", err)
	}
	
	s := echo.Server{}
	grpcServer := grpc.NewServer()

	echo.RegisterEchoServiceServer(grpcServer, &s)
	
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Serve error: ", err)
	}
	
}