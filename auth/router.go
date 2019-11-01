package auth

import (
	"YaIce/auth/mrg/inside"
	"YaIce/core/cluster"
	"YaIce/protobuf/internal_proto"
)

func registerRouter() {
	internal_proto.RegisterServiceConnectServer(cluster.Handler.GRpcServer, &inside.Service{})
}
