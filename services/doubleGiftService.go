package services

type IDoubleGiftService interface {
	getInfo() int
}
type DoubleGIftService struct{}

func (s DoubleGIftService) GetInfo() int {
	return 0
}

func (s DoubleGIftService) CheckHealth() bool {
	return true
}
