// Package product
package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ngdlong91/funtech/cmd/gin/dto"

	"github.com/gin-gonic/gin"
)

func Test_productService_Purchase(t *testing.T) {
	// Simulation http request with defined payload
	service := NewGinService()

	t.Run("invalid request", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Convert from request payload to io.Reader
		requestByte, _ := json.Marshal(dto.RequestPurchase{
			Id: 1,
			Products: []dto.Product{
				{
					Id:       5,
					Quantity: 5,
				},
				{
					Id:       1,
					Quantity: 1,
				},
			},
		})
		requestReader := bytes.NewReader(requestByte)
		c.Request, _ = http.NewRequest("POST", "/purchase", requestReader)

		service.Purchase(c)
		var respData dto.PurchaseResponse
		err := json.NewDecoder(w.Body).Decode(&respData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Data response %+v \n", respData)
		fmt.Println("Response code", w.Code)
	})

	t.Run("all product is success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Convert from request payload to io.Reader
		requestByte, _ := json.Marshal(dto.RequestPurchase{
			Id: 1,
			Products: []dto.Product{
				{
					Id:       5,
					Quantity: 5,
				},
				{
					Id:       1,
					Quantity: 1,
				},
			},
		})
		requestReader := bytes.NewReader(requestByte)
		c.Request, _ = http.NewRequest("POST", "/purchase", requestReader)

		service.Purchase(c)
		var respData dto.PurchaseResponse
		err := json.NewDecoder(w.Body).Decode(&respData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Data response %+v \n", respData)
		fmt.Println("Response code", w.Code)
	})

	t.Run("some of product success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Convert from request payload to io.Reader
		requestByte, _ := json.Marshal(dto.RequestPurchase{
			Id: 1,
			Products: []dto.Product{
				{
					Id:       5,
					Quantity: 5,
				},
				{
					Id:       1,
					Quantity: 1,
				},
			},
		})
		requestReader := bytes.NewReader(requestByte)
		c.Request, _ = http.NewRequest("POST", "/purchase", requestReader)

		service.Purchase(c)
		var respData dto.PurchaseResponse
		err := json.NewDecoder(w.Body).Decode(&respData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Printf("Data response %+v \n", respData)
		fmt.Println("Response code", w.Code)
	})
}
