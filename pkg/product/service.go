// Package product
package product

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/ngdlong91/funtech/dto"

	"github.com/ngdlong91/funtech/pkg/product/uc"
)

type Service interface {
	List()
	Detail()
	Purchase()
	Cancel()
}

type RPCHandler struct {
	log       *logrus.Entry
	productUC uc.ProductUseCase
}

func (s *RPCHandler) DoPurchase(c context.Context, payload *PurchaseRequest) (*PurchaseResponse, error) {
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
	//	Products: []dto.ProductUseCase{},
	//}
	//go s.productUC.Purchase(requestPayload)
	//s.productUC.Purchase(payload)

	s.log.Debugf("payload %+v \n", payload)

	internalPayload := dto.RequestPurchase{
		Id:       int(payload.CustomerId),
		Products: []dto.Product{},
	}

	for _, product := range payload.Products {
		internalProduct := dto.Product{
			Id:       int(product.Id),
			Quantity: int(product.Quantity),
		}
		internalPayload.Products = append(internalPayload.Products, internalProduct)
	}

	var response PurchaseResponse
	internalResponse, err := s.productUC.Purchase(internalPayload)
	if err != nil {
		return &PurchaseResponse{}, err
	}

	s.log.Debugf("Internal response %+v \n", internalResponse)
	for _, item := range internalResponse {
		result := PurchaseResult{
			Id:     int32(item.Id),
			Result: item.IsSuccess,
			Msg:    item.Result,
		}
		response.Result = append(response.Result, &result)
	}

	response.Response = &Response{
		Code: 200,
		Msg:  "Success",
	}

	s.log.Debugf("Final respsonse %+v \n", response)
	return &response, nil
}

func NewRPCHandler() *RPCHandler {
	return &RPCHandler{
		log:       logrus.WithField("rpcHandler", "product"),
		productUC: uc.NewProductProcessor(),
	}
}
