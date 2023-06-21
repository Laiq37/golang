package entities

type Product struct {
	ID                uint      `json:"id" gorm:"primary_key"`
	Name              string    `json:"name"`
	Stock             int       `json:"stock"`
	ProductCategoryID uint      `json:"categoryId"`
	OrderItem         OrderItem `json:"-"`
	CartItem          CartItem  `json:"-"`
}
