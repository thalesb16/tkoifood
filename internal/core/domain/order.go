package domain

type Order struct {
	ID      string `json:"id"`
	StoreID string `json:"store_id"`
	Order   string `json:"order"`
}

func NewOrder(orderID, storeID, order string) Order {
	return Order{
		ID:      orderID,
		StoreID: storeID,
		Order:   order,
	}
}
