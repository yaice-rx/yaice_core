package model

type MsgQueue struct {
	MsgNumber int32
	Token     []byte
	Session   *PlayerConn
	MsgData   []byte
}
