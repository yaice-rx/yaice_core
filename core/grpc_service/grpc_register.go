package grpc_service

import (
	"YaIce/protobuf/internal_proto"
	"google.golang.org/grpc"
)

func  RegisterServiceGrpc(server *grpc.Server) {
	internal_proto.RegisterServiceConnectServer(server,&ServiceRegister{})
	internal_proto.RegisterLoginVerifyServer(server,&LoginVerify{})
}

func RegisterClientGrpc(conn *grpc.ClientConn){
	internal_proto.NewServiceConnectClient(conn)
	internal_proto.NewLoginVerifyClient(conn)
}

