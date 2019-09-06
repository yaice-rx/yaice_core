package auth

import (
	"YaIce/auth/mrg/inside"
	"YaIce/core/handler"
	"YaIce/protobuf/internal_proto"
)

func registerRouter() {
	internal_proto.RegisterServiceConnectServer(handler.GRPCServer, &inside.Service{})
}
