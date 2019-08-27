package kcp_service

import "YaIce/core/model"

type MsgQueue struct {
	msgNumber int32
	Token     []byte
	Session   *model.PlayerConn
	msgData   []byte
}
