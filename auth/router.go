package auth

import (
	"YaIce/auth/mrg/inside"
	"YaIce/core/grpc_service"
	"YaIce/protobuf/internal_proto"
)

func registerRouter() {
	internal_proto.RegisterServiceConnectServer(grpc_service.GRpcServer, &inside.Service{})
}
