// Package dto
package dto

type RequestPurchase struct {
	Id       int       `json:"id"`
	Products []Product `json:"products"`
}
