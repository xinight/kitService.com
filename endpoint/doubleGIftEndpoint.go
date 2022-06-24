package endp

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"kittest.com/services"
)

type GetInfoResponse struct {
	GoldNum int `json:"gold_num"`
}

func GenGetInfoEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		res := services.DoubleGIftService{}.GetInfo()
		return GetInfoResponse{GoldNum: res}, nil
	}
}
