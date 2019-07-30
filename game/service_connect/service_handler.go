package service_connect

import (
	"YaIce/core/common"
	"YaIce/core/etcd_service"
	"YaIce/core/grpc_service"
	"YaIce/protobuf/internal_proto"
	"errors"
)

//组装客户端发送信息
func SendClientMsg(serverName string,msg interface{})error{
	if nil != etcd_service.EtcdClient.ConnServiceList[serverName]{
		return errors.New("server not exist")
	}
	var msgProtoNumber int32
	var msgData *internal_proto.MsgBodyRequest
	switch msg.(type) {
	case internal_proto.Request_ConnectStruct:
		msgData = &internal_proto.MsgBodyRequest{
			Connect: &internal_proto.Request_ConnectStruct{
			},
		}
		msgProtoNumber = common.ProtocalNumber(common.GetProtoName(&internal_proto.MsgBodyRequest{}))
		break;
	}
	data := &internal_proto.C_ServiceMsgRequest{
		MsgHandlerNumber:msgProtoNumber,
		Header:etcd_service.EtcdClient.Header,
		Struct:msgData,
	}
	return grpc_service.RegisterServiceRequest(*etcd_service.EtcdClient.ConnServiceList[serverName].Connect,data);
}


