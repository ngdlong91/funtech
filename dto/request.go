// Package dto
package dto

type RequestPurchase struct {
	Id       int       `json:"customer_id" binding:"required"`
	Products []Product `json:"products" binding:"required"`
}
