package grpc_service

import (
	"YaIce/protobuf/internal_proto"
	"context"
)

type ServiceRegister struct {

}

func (s *ServiceRegister)RegisterServiceRequest(  ctx context.Context,
	args *internal_proto.GameConnectServiceRequest)(response *internal_proto.GameConnectServiceReply, err error) {
	return &internal_proto.GameConnectServiceReply{},nil
}

//客户端连接
func RegisterServiceRequest(client internal_proto.ServiceConnectClient){
	var request internal_proto.GameConnectServiceRequest
	response, _ := client.RegisterServiceRequest(context.Background(), &request) //调用远程方法
	if response.IsConnected {
		
	}
}

type LoginVerify struct {

}

func (s *LoginVerify)LoginVerifyRequest(  ctx context.Context,
	args *internal_proto.GamePlayerLoginRequest)(response *internal_proto.GamePlayerLoginReply, err error) {
	return &internal_proto.GamePlayerLoginReply{
		Token:"-----------------",
	},nil
}

func LoginVerifyRequest( client internal_proto.LoginVerifyClient) {
	var request internal_proto.GamePlayerLoginRequest
	request.Guid = "++++++++++++++++";
	response, _ := client.LoginVerifyRequest(context.Background(), &request) //调用远程方法
	if("" != response.Token){

	}
}