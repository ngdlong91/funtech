// Package product
package product

import (
	"context"

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
}

type RPCHandler struct {
	productUC uc.Product
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

	ctx := context.Background()
	cancelCtx, cancel := context.WithCancel(ctx)
	requestPayload := dto.RequestPurchase{
		Id:       0,
		Products: []dto.Product{},
	}
	go s.productUC.Purchase(requestPayload)
}

func NewService() Service {

}
