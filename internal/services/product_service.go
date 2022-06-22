package services

import (
	"proj1/internal/dto"
	"proj1/internal/storage"
	"sync"
	"time"
)

var mutex sync.Mutex

func GetAllProducts() (responseProducts []dto.Product) {
	products := storage.Products
	for _, product := range products {
		if product.DeletedAt.IsZero() {
			responseProducts = append(responseProducts, product)
		}
	}
	return responseProducts
}

func GetProductByID(id uint) (responseProduct dto.Product) {
	products := storage.Products
	for _, product := range products {
		if product.Model.ID == uint(id) && product.DeletedAt.IsZero() {
			// if product found, the return product
			return product
		}
	}
	return
}

func InsertProduct(product dto.Product) (responseProduct dto.Product) {
	products := &storage.Products

	// use mutex to avoid race condition
	mutex.Lock()
	storage.ProductIndex = storage.ProductIndex + 1
	newID := storage.ProductIndex
	product.Model.ID = newID
	product.CreatedAt = time.Now()

	// add product to storage
	*products = append(*products, product)
	mutex.Unlock()
	return product
}

func UpdateProduct(id uint, product dto.Product) (responseProduct dto.Product) {
	products := &storage.Products
	for i := 0; i < len(*products); i++ {
		if (*products)[i].Model.ID == uint(id) && (*products)[i].DeletedAt.IsZero() {
			// if product found, then update the product
			(*products)[i].Name = product.Name
			(*products)[i].SKU = product.SKU
			(*products)[i].Price = product.Price
			(*products)[i].Quantity = product.Quantity
			(*products)[i].UpdatedAt = time.Now()
			product = (*products)[i]
			return product
		}
	}
	return
}

func DeleteProduct(id uint) (responseProduct dto.Product) {
	products := &storage.Products
	for i := 0; i < len(*products); i++ {
		if (*products)[i].Model.ID == uint(id) && (*products)[i].DeletedAt.IsZero() {
			// if product found, then add time to DeletedAt field (soft delete)
			(*products)[i].DeletedAt = time.Now()
			responseProduct = (*products)[i]
			return
		}
	}
	return
}
