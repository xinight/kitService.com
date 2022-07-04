package services

type IDoubleGiftService interface {
	getInfo() int
}
type DoubleGIftService struct{}

const (
	EXCHANGE_SUCCESS = iota
	ALREADY_EXCHANGED
	INTEGRAL_NOT_ENOUGH
	UNDEFINDED_INDEX
)

var ExchangeStatus = [4]string{"", "goods already exchange", "integral not enough", "undefinded index"}

func (s DoubleGIftService) CheckHealth() bool {
	return true
}

func (s DoubleGIftService) GetInfo() int {
	return 0
}

func (s DoubleGIftService) Exchange(idx int) int {
	return EXCHANGE_SUCCESS
}
