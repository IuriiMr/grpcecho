package main

import (
	"context"
	"log"

	"github.com/iuriimr/grpcecho/echo"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Instaed of Postaman, we can use the following command to test the server:
func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Dial error: ", err)
	}
	defer conn.Close()

	c := echo.NewEchoServiceClient(conn)
	msg := echo.Message{Body: "Hello from client!"}

	response, err := c.ReturnEcho(context.Background(), &msg)
	if err != nil {
		log.Fatal("Response error: ", err)
	}
	log.Printf("Response from server: %s", response.Body)
}