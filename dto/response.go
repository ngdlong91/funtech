// Package dto
package dto

/*
Response define internal response for tracking and display.
Now I use 3 code to define where we should invest time
*/

const (
	CodeClientErr   = 1
	CodeProcessErr  = 2
	CodeExternalErr = 3
)

type Response struct {
	Code int
	Msg  string
}

type PurchaseResponse struct {
	Response
	Results []PurchaseResult `json:"results"`
}

type PurchaseResult struct {
	Id        int  `json:"id"`
	IsSuccess bool `json:"is_success"`
}
