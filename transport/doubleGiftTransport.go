package mtsp

import (
	"context"
	"encoding/json"
	"net/http"
)

func GetInfoDecodeRequest(c context.Context, w *http.Request) (request interface{}, err error) {
	return nil, nil
}

func GetInfoEncodeResponse(c context.Context, w http.ResponseWriter, respose interface{}) error {
	return json.NewEncoder(w).Encode(respose)
}
