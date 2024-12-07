package order

import "github.com/sjxiang/oms/order/entity"


type Order struct {
	ID          string
	CustomerID  string
	Status      string
	PaymentLink string
	Items       []*entity.Item
}


func NewOrder(id, customerID, status, paymentLink string, items []*entity.Item) (*Order, error) {
	return &Order{}, nil
}


func NewPendingOrder() (*Order, error) {
	return nil, nil
}

func (o *Order) IsPaid() error {
	return nil
}


// 是否已经支付
func (o *Order) HasPaid() bool {
	return false
}