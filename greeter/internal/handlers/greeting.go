package handlers

import (
	"context"
	"fmt"

	"greeter/internal/gen"
)

// Server implements gen.StrictServerInterface.
type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetGreeting(_ context.Context, request gen.GetGreetingRequestObject) (gen.GetGreetingResponseObject, error) {
	if request.Params.Name == "" {
		return gen.GetGreeting400JSONResponse{
			Code:    400,
			Message: "name is required",
		}, nil
	}

	return gen.GetGreeting200JSONResponse{
		Message: fmt.Sprintf("Hello, %s!", request.Params.Name),
	}, nil
}
