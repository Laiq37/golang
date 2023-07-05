package entities

type Order struct {
	ID         uint `json:"id" gorm:"primary_key"`
	UserID     uint
	OrderTime  int64       `json:"orderTime"`
	OrderItems []OrderItem `json:"orderItems"`
	Payment    Payment     `json:"payment"`
	Status     string      `json:"status" gorm:"default:'pending'"`
}
