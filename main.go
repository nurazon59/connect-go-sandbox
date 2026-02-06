package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	hellov1 "github.com/nurazon59/connect-go-sandbox/gen/hello/v1"
	"github.com/nurazon59/connect-go-sandbox/gen/hello/v1/hellov1connect"
)

type HelloServer struct{}

func (s *HelloServer) Hello(
	ctx context.Context,
	req *connect.Request[hellov1.HelloRequest],
) (*connect.Response[hellov1.HelloResponse], error) {
	body := req.Msg.Str
	return connect.NewResponse(&hellov1.HelloResponse{
		Message: fmt.Sprintf("Hello ,%s", body),
	}), nil
}

func main() {
	mux := http.NewServeMux()

	path, handler := hellov1connect.NewHelloServiceHandler(&HelloServer{})
	mux.Handle(path, handler)

	http.ListenAndServe("localhost:8080", mux)
}
