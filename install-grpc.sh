#下载grpc
git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text

#自行下载安装protoc，下载地址https://github.com/google/protobuf/releases

#下载protoc插件
go get -u github.com/golang/protobuf/protoc-gen-go
