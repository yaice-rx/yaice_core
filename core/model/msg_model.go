package model

type MsgQueue struct {
	MsgNumber int32
	Session   *Conn
	MsgData   []byte
}
