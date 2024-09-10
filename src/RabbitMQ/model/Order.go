package rabbitMQ

type Order struct {
	CustomerId    string `json:"customerId"`
	CustomerName  string `json:"customerName"`
	OrderDateTime string `json:"orderDateTime"`
	TotalAmount   float64 `json:"totalAmount"`
	TotalCurrency string `json:"totalCurrency"`
	OrderItems     []interface{}    `json:"orderItems"`
}

type OrderItem struct{
	productId  string `json:"productId"`
	quantity   int     `json:"quantity"`
	amount     float64  `json:"amount"`
	currency   string   `json:"currency"`
}
