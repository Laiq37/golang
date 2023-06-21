package entities

type ProductCategory struct {
	ID       uint      `json:"cateogoryId" gorm:"primary_key"`
	Name     string    `json:"name" gorm:"unique"`
	Products []Product `json:"products" gorm:"foreignKey:ProductCategoryID"`
}
