package dtsp

import (
	"context"

	kgrpc "github.com/go-kit/kit/transport/grpc"
	"kittest.com/def"
	endp "kittest.com/endpoint"
	pb "kittest.com/pbs"
	"kittest.com/util"
)

type doubleGiftRpc struct {
	GetInfoHandler kgrpc.Handler
	pb.UnimplementedGetInfoServer
	ExchangeHandler kgrpc.Handler
	pb.UnimplementedExchangeServer
}

func GetInfoRpcDecodeRequest(c context.Context, gprcReq interface{}) (interface{}, error) {
	_, ok := gprcReq.(*pb.GetInfoReq)
	if !ok {
		return util.Response(true, "GetInfoRpcDecodeRequest 入参断言错误", nil)
	}
	return &pb.GetInfoReq{}, nil
}

func GetInfoRpcEncodeResponse(c context.Context, res interface{}) (interface{}, error) {
	r, ok := res.(def.GetInfoResponse)
	if !ok {
		return util.Response(true, "GetInfoRpcEncodeResPonse 返回断言错误", nil)
	}
	response := &pb.GetInfoRes{GoldNum: int32(r.GoldNum)}

	return response, nil
}

func ExchangeRpcDecodeRequest(c context.Context, gprcReq interface{}) (interface{}, error) {
	r, ok := gprcReq.(*pb.ExchangeReq)
	if !ok {
		return util.Response(true, "ExchangeRpcDecodeRequest 入参断言错误", nil)
	}
	req := def.ExchangeRequest{
		Index: int(r.Idx),
	}
	return req, nil
}

func ExchangeRpcEncodeResponse(c context.Context, res interface{}) (interface{}, error) {
	r, ok := res.(def.ExchangeRqsponse)
	if !ok {
		return util.Response(true, "ExchangeRpcEncodeResponse 返回断言错误", nil)
	}
	response := &pb.ExchangeRes{Gotten: int32(r.Gotten)}
	return response, nil
}

func (d *doubleGiftRpc) GetInfoRpc(c context.Context, req *pb.GetInfoReq) (*pb.GetInfoRes, error) {
	_, res, err := d.GetInfoHandler.ServeGRPC(c, req)
	if err != nil {
		_, e := util.Response(true, err.Error(), nil)
		return nil, e
	}
	return res.(*pb.GetInfoRes), nil
}

func (d *doubleGiftRpc) ExchangeRpc(c context.Context, req *pb.ExchangeReq) (*pb.ExchangeRes, error) {
	_, res, err := d.ExchangeHandler.ServeGRPC(c, req)
	if err != nil {
		_, e := util.Response(true, err.Error(), nil)
		return nil, e
	}
	return res.(*pb.ExchangeRes), nil
}

func NewDoubleGiftRpc() *doubleGiftRpc {
	d := new(doubleGiftRpc)
	d.GetInfoHandler = kgrpc.NewServer(
		endp.GenGetInfoEndpoint(),
		GetInfoRpcDecodeRequest,
		GetInfoRpcEncodeResponse,
	)
	d.ExchangeHandler = kgrpc.NewServer(
		endp.GenExchangeEndpoint(),
		ExchangeRpcDecodeRequest,
		ExchangeRpcEncodeResponse,
	)

	return d

}
