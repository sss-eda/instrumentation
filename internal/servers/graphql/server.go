package graphql

import (
	"fmt"
	"time"
)

// Server TODO
type Server struct{}

// NewServer TODO
func NewServer() (*Server, error) {
	return &Server{}, nil
}

func (server *Server) Run() {
	for {
		fmt.Println("GraphQL Server is running...")
		time.Sleep(time.Second)
	}
}
