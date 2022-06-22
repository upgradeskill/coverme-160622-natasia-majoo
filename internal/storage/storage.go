package storage

import (
	"proj1/internal/dto"
	"time"
)

var Products []dto.Product
var ProductIndex uint

func Seeder() {
	Products = []dto.Product{
		{Model: dto.Model{ID: 1, CreatedAt: time.Now()}, Name: "Coffee", SKU: "PR-001", Price: 10000, Quantity: 20},
		{Model: dto.Model{ID: 2, CreatedAt: time.Now()}, Name: "Milk", SKU: "PR-002", Price: 6000, Quantity: 50},
	}
	ProductIndex = 2
}
