package entities

type CartItem struct {
	UserID    uint `json:"userId" gorm:"primaryKey;autoIncrement:false"`
	ProductID uint `json:"productId" gorm:"primaryKey;autoIncrement:false"`
	Quantity  uint `json:"quantity"`
}
