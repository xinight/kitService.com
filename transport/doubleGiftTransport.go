package dtsp

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"kittest.com/def"
	"kittest.com/util"
)

func GetInfoDecodeRequest(c context.Context, r *http.Request) (request interface{}, err error) {
	return nil, nil
}

func ExchangeDecodeRequest(c context.Context, r *http.Request) (request interface{}, err error) {
	if r.URL.Query().Get("idx") != "" {
		idx, _ := strconv.Atoi(r.URL.Query().Get("idx"))
		return def.ExchangeRequest{Index: idx}, nil
	}
	return util.Response(true, "参数错误", nil)
}

func EncodeResponse(c context.Context, w http.ResponseWriter, respose interface{}) error {
	w.Header().Set("Content-type", "application/json")
	res, err := util.Response(false, "success", respose)
	json.NewEncoder(w).Encode(res)
	return err
}
