package common

import (
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"strings"
	"sync"
)

var mutex sync.Mutex

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

//把协议名称转为唯一协议编号
func ProtocalNumber(replacement string)int{
	var h int32
	h = 0
	for _, char := range []rune(replacement) {
		h = 31 * h + int32(char)
	}
	return int(h)
}

//获取协议名称
func GetProtoName(t proto.Message)string{
	x := proto.MessageName(t)
	proto_ :=strings.Split(x,".")
	if len(proto_) > 0 {
		return proto_[1]
	}else{
		return ""
	}
}

//连个字符串的key合并
func MergeMapString(varA map[string]string,varB  map[string]string)map[string]string{
	data := make(map[string]string, len(varA)+len(varB))
	for k,v := range varA{
		data[k] = v
	}
	for k,v := range varB{
		data[k] = v
	}
	return data
}
