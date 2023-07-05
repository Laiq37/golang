package entities

type OrderItem struct {
	OrderID   uint   `json:"-"`
	Quantity  int    `json:"quantity"`
	ProductID uint   `json:"productId"`
	Name      string `json:"productName"`
}
