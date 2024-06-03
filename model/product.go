package model

type Product struct {
	ProductID   int64  `gorm:"primaryKey" json:"product_id"`
	ProductName string `gorm:"type:varchar(200)" json:"product_name"`
	Description string `gorm:"type:text" json:"description"`
}
