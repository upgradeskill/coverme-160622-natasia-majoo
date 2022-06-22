package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"proj1/internal/dto"
	"proj1/internal/handler"
	"proj1/internal/storage"
	"strings"
	"testing"
)

// run `go test -cover`` in this path
func TestGetAllProducts(t *testing.T) {
	// call seeder for data dummy
	storage.Seeder()
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.GetAllProducts())
	handler.ServeHTTP(w, req)

	expectedResult := strings.TrimSpace(`[{"id":1,"created_at":"2022-06-20T23:05:23.5272439+07:00","updated_at":"0001-01-01T00:00:00Z","deleted_at":"0001-01-01T00:00:00Z","name":"Coffee","sku":"PR-001","price":10000,"qty":20},{"id":2,"created_at":"2022-06-20T23:05:23.5272439+07:00","updated_at":"0001-01-01T00:00:00Z","deleted_at":"0001-01-01T00:00:00Z","name":"Milk","sku":"PR-002","price":6000,"qty":50}]`)
	var expedtedResultObject []dto.Product
	var resultObject []dto.Product

	// if status code is not 200, then test failed
	if status := w.Code; status != http.StatusOK {
		t.Errorf("status code not success: got %v", status)
	}

	if err := json.Unmarshal([]byte(expectedResult), &expedtedResultObject); err != nil {
		t.Errorf("test failed: failed to unmarshal string to object")
	}

	if err := json.Unmarshal(w.Body.Bytes(), &resultObject); err != nil {
		t.Errorf("test failed: failed to unmarshal string to object")
	}

	for i := 0; i < len(expedtedResultObject); i++ {
		// compare for every data in expected and result, exclude created_at, updated_at, and deleted_at (as it get from time.Now)
		if expedtedResultObject[i].ID != resultObject[i].ID || expedtedResultObject[i].Name != resultObject[i].Name || expedtedResultObject[i].SKU != resultObject[i].SKU || expedtedResultObject[i].Price != resultObject[i].Price || expedtedResultObject[i].Quantity != resultObject[i].Quantity {
			t.Errorf("test result mismatched with expected result")
			return
		}
	}

}

func TestGetProductByID(t *testing.T) {
	// call seeder for data dummy
	storage.Seeder()

	req, _ := http.NewRequest(http.MethodGet, "/product?id=1", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.ProductHandler())
	handler.ServeHTTP(w, req)

	expectedResult := `{"id":1,"created_at":"2022-06-20T22:54:08.3234809+07:00","updated_at":"0001-01-01T00:00:00Z","deleted_at":"0001-01-01T00:00:00Z","name":"Coffee","sku":"PR-001","price":10000,"qty":20}`

	// if status code is not 200, then test failed
	if status := w.Code; status != http.StatusOK {
		t.Errorf("test failed: status code not success, got %v", status)
	}

	var expedtedResultObject dto.Product
	var resultObject dto.Product

	if err := json.Unmarshal([]byte(expectedResult), &expedtedResultObject); err != nil {
		t.Errorf("test failed: failed to unmarshal string to object")
	}

	if err := json.Unmarshal(w.Body.Bytes(), &resultObject); err != nil {
		t.Errorf("test failed: failed to unmarshal string to object")
	}

	// compare for result and expected result, except the timestamp (created_at, updated_at, deleted_at)
	if expedtedResultObject.ID != resultObject.ID || expedtedResultObject.Name != resultObject.Name || expedtedResultObject.SKU != resultObject.SKU || expedtedResultObject.Price != resultObject.Price || expedtedResultObject.Quantity != resultObject.Quantity {
		t.Errorf("test failed:  mismatched with expected result")
		return
	}
}

func TestGetProductByID_WhenProductNotFound(t *testing.T) {
	// call seeder for data dummy
	storage.Seeder()

	req, _ := http.NewRequest(http.MethodGet, "/product?id=10", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.ProductHandler())
	handler.ServeHTTP(w, req)

	// if status code is not 400, then test failed
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("test failed: status code not success, got %v", status)
	}
}

func TestInsertProduct(t *testing.T) {
	// call seeder for data dummy
	storage.Seeder()
	// data to insert
	bodyReq := strings.NewReader(`{"name": "Tea","sku": "PR-003","price": 3000,"qty": 50}`)
	req, _ := http.NewRequest(http.MethodPost, "/product", bodyReq)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.ProductHandler())
	handler.ServeHTTP(w, req)

	expectedResult := `{"id":3,"name":"Tea","sku":"PR-003","price":3000,"qty":50}`

	// if status code not 200, test failed
	if status := w.Code; status != http.StatusOK {
		t.Errorf("test failed: status code not success, got %v", status)
	}

	var expedtedResultObject dto.Product
	var resultObject dto.Product

	if err := json.Unmarshal([]byte(expectedResult), &expedtedResultObject); err != nil {
		t.Errorf("test failed: failed to unmarshal string to object")
	}

	if err := json.Unmarshal(w.Body.Bytes(), &resultObject); err != nil {
		t.Errorf("test failed: failed to unmarshal string to object")
	}

	// compare for result and expected result, except the timestamp (created_at, updated_at, deleted_at)
	if expedtedResultObject.ID != resultObject.ID || expedtedResultObject.Name != resultObject.Name || expedtedResultObject.SKU != resultObject.SKU || expedtedResultObject.Price != resultObject.Price || expedtedResultObject.Quantity != resultObject.Quantity {
		t.Errorf("test failed:  mismatched with expected result")
		return
	}
}

func TestInsertProduct_WhenWrongBodyRequest(t *testing.T) {
	// call seeder for data dummy
	storage.Seeder()
	// data to insert
	bodyReq := strings.NewReader(`{"name": "Tea","sku": "PR-003","price": "test","qty": "test"}`)
	req, _ := http.NewRequest(http.MethodPost, "/product", bodyReq)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.ProductHandler())
	handler.ServeHTTP(w, req)

	// if status code not 400, test failed
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("test failed: status code not success, got %v", status)
	}
}

func TestUpdateProduct(t *testing.T) {
	// call seeder for data dummy
	storage.Seeder()
	// data to update
	bodyReq := strings.NewReader(`{"name": "Tea","sku": "PR-003","price": 3000,"qty": 50}`)
	// update from product id 2
	req, _ := http.NewRequest(http.MethodPut, "/product?id=2", bodyReq)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.ProductHandler())
	handler.ServeHTTP(w, req)

	expectedResult := `{"id":2,"name":"Tea","sku":"PR-003","price":3000,"qty":50}`

	// if status code not 200, test failed
	if status := w.Code; status != http.StatusOK {
		t.Errorf("test failed: status code not success, got %v", status)
	}

	var expedtedResultObject dto.Product
	var resultObject dto.Product

	if err := json.Unmarshal([]byte(expectedResult), &expedtedResultObject); err != nil {
		t.Errorf("test failed: failed to unmarshal string to object")
	}

	if err := json.Unmarshal(w.Body.Bytes(), &resultObject); err != nil {
		t.Errorf("test failed: failed to unmarshal string to object")
	}

	// compare for result and expected result, except the timestamp (created_at, updated_at, deleted_at)
	if expedtedResultObject.ID != resultObject.ID || expedtedResultObject.Name != resultObject.Name || expedtedResultObject.SKU != resultObject.SKU || expedtedResultObject.Price != resultObject.Price || expedtedResultObject.Quantity != resultObject.Quantity {
		t.Errorf("test failed:  mismatched with expected result")
		return
	}
}

func TestUpdateProduct_WhenWrongBodyRequest(t *testing.T) {
	// call seeder for data dummy
	storage.Seeder()
	// data to update
	bodyReq := strings.NewReader(`{"name": "Tea","sku": "PR-003","price": "test","qty": "test"}`)
	// update from product id 2
	req, _ := http.NewRequest(http.MethodPut, "/product?id=20", bodyReq)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.ProductHandler())
	handler.ServeHTTP(w, req)

	// if status code not 400, test failed
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("test failed: status code not success, got %v", status)
	}
}

func TestUpdateProduct_WhenIdNotFound(t *testing.T) {
	// call seeder for data dummy
	storage.Seeder()
	// data to update
	bodyReq := strings.NewReader(`{"name": "Tea","sku": "PR-003","price": 3000,"qty": 50}`)
	// update from product id 2
	req, _ := http.NewRequest(http.MethodPut, "/product?id=20", bodyReq)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.ProductHandler())
	handler.ServeHTTP(w, req)

	// if status code not 400, test failed
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("test failed: status code not success, got %v", status)
	}
}

func TestDeleteProduct(t *testing.T) {
	// call seeder for data dummy
	storage.Seeder()

	req, _ := http.NewRequest(http.MethodDelete, "/product?id=2", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.ProductHandler())
	handler.ServeHTTP(w, req)

	expectedResult := `{"id":2,"name":"Milk","sku":"PR-002","price":6000,"qty":50}`

	if status := w.Code; status != http.StatusOK {
		t.Errorf("test failed: status code not success, got %v", status)
	}

	var expedtedResultObject dto.Product
	var resultObject dto.Product

	if err := json.Unmarshal([]byte(expectedResult), &expedtedResultObject); err != nil {
		t.Errorf("test failed: failed to unmarshal string to object")
	}

	if err := json.Unmarshal(w.Body.Bytes(), &resultObject); err != nil {
		t.Errorf("test failed: failed to unmarshal string to object")
	}

	// compare for result and expected result, except the timestamp (created_at, updated_at, deleted_at)
	if expedtedResultObject.ID != resultObject.ID || expedtedResultObject.Name != resultObject.Name || expedtedResultObject.SKU != resultObject.SKU || expedtedResultObject.Price != resultObject.Price || expedtedResultObject.Quantity != resultObject.Quantity {
		t.Errorf("test failed: mismatched with expected result")
		return
	}

	// if deleted_at is null, then the soft delete is not working (test failed)
	if resultObject.DeletedAt.IsZero() {
		t.Errorf("test failed: failed to delete data")
		return
	}

}

func TestDeleteProduct_WhenProductNotFound(t *testing.T) {
	// call seeder for data dummy
	storage.Seeder()

	req, _ := http.NewRequest(http.MethodDelete, "/product?id=10", nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.ProductHandler())
	handler.ServeHTTP(w, req)

	// if status code is not 400, then test failed
	if status := w.Code; status != http.StatusBadRequest {
		t.Errorf("test failed: status code not success, got %v", status)
	}
}
