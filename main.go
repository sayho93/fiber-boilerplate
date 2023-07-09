package main

import (
	"context"
	"fiber/src"
	proto "fiber/src/proto"
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"sync"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

type gRpcServer struct {
	proto.UnimplementedCalcServiceServer
}

func runHttpServ(address string, port string) {
	server, _ := src.New()
	log.Info(fmt.Sprintf("HTTP Listening on port %s", port))
	err := server.Listen(fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		panic(err)
	}
}

func runGRPCServ(address string, port string) {
	log.Info(fmt.Sprintf("TCP Listening on port %s", port))
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", address, port))

	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterCalcServiceServer(srv, &gRpcServer{})
	reflection.Register(srv)
	if e := srv.Serve(lis); e != nil {
		panic(err)
	}

	log.Info("test222")
}

func main() {
	port := os.Getenv("PORT")
	address := func(appEnv string) string {
		if appEnv == "development" {
			return "localhost"
		}
		return ""
	}(os.Getenv("APP_ENV"))

	var wg sync.WaitGroup
	wg.Add(2)
	go runHttpServ(address, port)
	go runGRPCServ(address, "4040")
	wg.Wait()
}

func (s *gRpcServer) Add(_ context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b

	return &proto.Response{Result: result}, nil
}

func (s *gRpcServer) Multiply(_ context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a * b

	return &proto.Response{Result: result}, nil
}
