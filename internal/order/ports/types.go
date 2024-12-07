package ports

type CreateOrderRequest struct {
	CustomerID  string              `json:"customer_id"`
	Items       []ItemWithQuantity  `json:"items"`
}

type Item struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	PriceId  string `json:"price_id"`
	Quantity int32  `json:"quantity"`
}

type ItemWithQuantity struct {
	Id       string `json:"id"`
	Quantity int32  `json:"quantity"`
}

type Order struct {
	CustomerID  string    `json:"customer_id"`
	Id          string    `json:"id"`
	Items       []Item    `json:"items"`
	PaymentLink string    `json:"payment_link"`
	Status      string    `json:"status"`
}

type Response struct {
	Data    map[string]interface{} `json:"data"`
	Errno   int32                  `json:"errno"`
	Message string                 `json:"message"`
	TraceID string                 `json:"trace_id"`
}