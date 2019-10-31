// Package dto
package dto

type Response struct {
	Code int
	Msg  string
}

type PurchaseResponse struct {
	Response
}

type PurchaseResult struct {
	Id        int  `json:"id"`
	IsSuccess bool `json:"is_success"`
	Remain    int  `json:"remain"` // Remain quantity in db
}
