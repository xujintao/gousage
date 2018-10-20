package main

import (
	"context"
	"log"
	"os"
	pb "test/testgrpc/service"
	"time"

	"google.golang.org/grpc"
)

func main() {
	log.SetOutput(os.Stdout)
	conn, err := grpc.Dial("192.168.6.201:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//矩形
	rect := pb.NewRectClient(conn)
	r, err := rect.Area(ctx, &pb.RectParamIn{Width: 50, Height: 100})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r.Result)

	r, err = rect.Perimeter(ctx, &pb.RectParamIn{Width: 50, Height: 100})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r.Result)

	//鉴权
	auth := pb.NewAuthClient(conn)
	r2, err := auth.Auth(ctx, &pb.AuthParamIn{Uid: 123, TokenId: "abcd123"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r2.Result)
}
