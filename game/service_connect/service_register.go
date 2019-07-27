package service_connect

import (
	"YaIce/protobuf/internal_proto"
	"google.golang.org/grpc"
)


func RegisterClientGrpc(conn *grpc.ClientConn){
	internal_proto.NewServiceConnectClient(conn)
	internal_proto.NewLoginVerifyClient(conn)
}

