package internal

import (
	"fmt"
	"net/http"
	"proj1/internal/handler"
)

var baseURL = "http://localhost:8000"

func DoesSomethingAndReturn5() uint {
	fmt.Println("I did something")
	return 5
}

func SetRoutes() {

	// routing for get all product
	http.HandleFunc("/products", handler.GetAllProducts())

	// group routing for product
	http.HandleFunc("/product", handler.ProductHandler())

	fmt.Println("Listening in Server: " + baseURL)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
