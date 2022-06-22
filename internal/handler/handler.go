package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"proj1/internal/dto"
	"proj1/internal/services"
	"strconv"
)

func GetAllProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET /products")
		products := services.GetAllProducts()

		err := json.NewEncoder(w).Encode(products)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err) //nolint
		}
	}
}

func GetProductByID(w http.ResponseWriter, r *http.Request) (product dto.Product, response dto.ResponseError) {
	// to get and convert parameter into int (id)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return product, dto.ResponseError{Error: err}
	}
	fmt.Printf("GET /product/%d\n", id)
	product = services.GetProductByID(uint(id))
	// if product found
	if product.ID != 0 {
		return
	}
	// if product not found, then return error
	w.WriteHeader(http.StatusBadRequest)
	return product, dto.ResponseError{Message: "failed to get data: product not found"}
}

func InsertProduct(w http.ResponseWriter, r *http.Request) (product dto.Product, response dto.ResponseError) {
	fmt.Println("POST /product")

	// get data from request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	// parse reqBody into product object
	err := json.Unmarshal(reqBody, &product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return product, dto.ResponseError{Error: err}
	}

	product = services.InsertProduct(product)
	if product.ID != 0 {
		return
	}

	return
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) (product dto.Product, response dto.ResponseError) {
	// to get and convert parameter into int (id)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return product, dto.ResponseError{Error: err}
	}
	fmt.Printf("PUT /product/%d\n", id)

	// get data from request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	// parse reqBody into product object
	err = json.Unmarshal(reqBody, &product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return product, dto.ResponseError{Error: err}
	}

	product = services.UpdateProduct(uint(id), product)
	if product.ID != 0 {
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	return product, dto.ResponseError{Message: "failed to update: product not found"}
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) (product dto.Product, response dto.ResponseError) {
	// to get and convert parameter into int (id)
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return product, dto.ResponseError{Error: err}
	}
	fmt.Printf("DELETE /product/%d\n", id)

	product = services.DeleteProduct(uint(id))
	if product.ID != 0 {
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	return product, dto.ResponseError{Message: "failed to delete: product not found"}
}

func ProductHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// call function based in method
		switch r.Method {
		case http.MethodGet:
			product, response := GetProductByID(w, r)
			if response.Error != nil || response.Message != "" {
				json.NewEncoder(w).Encode(response) //nolint
				return
			}
			json.NewEncoder(w).Encode(product) //nolint
			return

		case http.MethodPost:
			product, response := InsertProduct(w, r)
			if response.Error != nil || response.Message != "" {
				json.NewEncoder(w).Encode(response) //nolint
				return
			}
			json.NewEncoder(w).Encode(product) //nolint
			return

		case http.MethodPut:
			product, response := UpdateProduct(w, r)
			if response.Error != nil || response.Message != "" {
				json.NewEncoder(w).Encode(response) //nolint
				return
			}
			json.NewEncoder(w).Encode(product) //nolint
			return

		case http.MethodDelete:
			product, response := DeleteProduct(w, r)
			if response.Error != nil || response.Message != "" {
				json.NewEncoder(w).Encode(response) //nolint
				return
			}
			json.NewEncoder(w).Encode(product) //nolint
			return
		}
	}
}
