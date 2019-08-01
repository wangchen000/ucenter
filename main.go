package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"ucenter/rpc"
	"ucenter/service"
)

//go:generate protoc -I proto/ proto/user.proto --go_out=plugins=grpc:./rpc
func main() {

	lis ,err := net.Listen("tcp",fmt.Sprintf(":%d", 1368))
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	rpc.RegisterUserServiceServer(server, &service.UserService{})

	err = server.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
