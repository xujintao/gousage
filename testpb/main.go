package main

import (
	"fmt"
	pb "test/testpb/service"

	"github.com/golang/protobuf/proto"
)

func main() {
	in := &pb.AuthParamIn{Uid: 123, TokenId: "abcd"}

	dataIn, _ := proto.Marshal(in)
	fmt.Println(string(dataIn))

	out := &pb.AuthParamIn{}
	proto.Unmarshal(dataIn, out)
	fmt.Println(out)
}
