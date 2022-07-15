package endp

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"kittest.com/def"
	"kittest.com/services"
	"kittest.com/util"
)

func GenGetInfoEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res := services.DoubleGIftService{}.GetInfo()
		response = def.GetInfoResponse{GoldNum: res}
		return response, nil
	}
}

func GenExchangeEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(def.ExchangeRequest)
		res := services.DoubleGIftService{}.Exchange(req.Index)
		if res > services.EXCHANGE_SUCCESS {
			return util.Response(true, services.ExchangeStatus[res], def.ExchangeRqsponse{})
		}
		response = def.ExchangeRqsponse{Gotten: res}
		return
	}
}
