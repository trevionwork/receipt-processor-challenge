package models

type Receipt struct {
	Retailer     string        `json:"retailer" binding:"required,retailer"`
	PurchaseDate string        `json:"purchaseDate" binding:"required,dateString"`
	PurchaseTime string        `json:"purchaseTime" binding:"required,timeString"`
	Items        []ReceiptItem `json:"items" binding:"required,min=1,dive"`
	Total        string        `json:"total" binding:"required,currencyString"`
}
