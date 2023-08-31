package echo

import (
	"log"
	"context"
)

type Server struct {
	UnimplementedEchoServiceServer
}

func (s *Server) ReturnEcho(ctx context.Context, msg *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", msg.Body)
	// return &Message{Body: "Hello from server!"}, nil
	return msg, nil
}

