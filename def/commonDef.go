package def

type FormatResponse struct {
	Err  bool        `json:"err"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type GetInfoResponse struct {
	GoldNum int `json:"gold_num"`
}

type ExchangeRequest struct {
	Index int `json:"index"`
}
type ExchangeRqsponse struct {
	Gotten int `json:"gotten"`
}
