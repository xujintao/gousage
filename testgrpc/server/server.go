package main

import (
	"context"
	"log"
	"net"
	"os"
	pb "test/testgrpc/service"

	"google.golang.org/grpc"
)

type myServer struct{}

func (s *myServer) Area(ctx context.Context, in *pb.RectParamIn) (*pb.RectParamOut, error) {
	ret := in.Width * in.Height
	return &pb.RectParamOut{Result: ret}, nil
}

func (s *myServer) Perimeter(ctx context.Context, in *pb.RectParamIn) (*pb.RectParamOut, error) {
	ret := in.Width + in.Height
	return &pb.RectParamOut{Result: ret}, nil
}

func (s *myServer) Auth(ctx context.Context, in *pb.AuthParamIn) (*pb.AuthParamOut, error) {
	//查询redis
	var ret string
	if in.Uid == 123 && in.TokenId == "abcd123" {
		ret = "success"
	} else {
		ret = "failed"
	}
	return &pb.AuthParamOut{Result: ret}, nil
}

func main() {
	log.SetOutput(os.Stdout)

	s := grpc.NewServer()
	// reflection.Register(s)

	//注册矩形服务
	pb.RegisterRectServer(s, &myServer{})
	//注册鉴权服务
	pb.RegisterAuthServer(s, &myServer{})

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
