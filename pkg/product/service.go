// Package product
package product

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ngdlong91/funtech/cmd/gin/dto"
	"github.com/ngdlong91/funtech/cmd/gin/pkg/product/uc"
)

type Service interface {
	List()
	Detail()
	Purchase()
	Cancel()
}

type GinHandler struct {
	processor uc.Product
}

func (s *GinHandler) Purchase(c *gin.Context) {
	var requestPayload dto.RequestPurchase
	if err := c.ShouldBindJSON(&requestPayload); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Code: dto.CodeClientErr,
			Msg:  "payload invalid",
		})
		return
	}

	if err := s.processor.Purchase(requestPayload); err != nil {
		c.JSON(http.StatusServiceUnavailable, dto.Response{
			Code: dto.CodeProcessErr,
			Msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.PurchaseResponse{
		Response: dto.Response{
			Msg: "success",
		},
		Results: []dto.PurchaseResult{},
	})

}

type RPCHandler struct {
	productUC uc.Product
}

func (s *RPCHandler) DoPurchase(context.Context, *PurchaseRequest) (*PurchaseResponse, error) {
	return &PurchaseResponse{}, nil
}

func (s *RPCHandler) Purchase() {
	// Problems:
	// We have 100 req at the same time and 20/100 req should success with some item
	// So the other should fail and cancel the current process
	// Fail reason:
	// 1. Remain stock is not enough >> cancel all remain request (we can use the in-memory storage
	// for quick check before update database)
	// 2. User may not finish the transaction (not enough balance, cancel, system error...)
	// 3. The stock refill
	// all transaction should use same parent context

	//ctx := context.Background()
	//cancelCtx, cancel := context.WithCancel(ctx)
	//requestPayload := dto.RequestPurchase{
	//	Id:       0,
	//	Products: []dto.Product{},
	//}
	//go s.productUC.Purchase(requestPayload)
}

func NewGinService() *GinHandler {
	return &GinHandler{
		processor: uc.NewProductProcessor(),
	}
}

func NewRPCHandler() *RPCHandler {
	return &RPCHandler{}
}
