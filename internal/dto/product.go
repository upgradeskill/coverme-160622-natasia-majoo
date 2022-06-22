package dto

type Product struct {
	Model
	Name     string `json:"name" form:"name"`
	SKU      string `json:"sku" form:"sku"`
	Price    uint   `json:"price" form:"price"`
	Quantity uint   `json:"qty" form:"qty"`
}
