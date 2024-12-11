package models

type ReceiptItem struct {
	ShortDescription string `json:"shortDescription" binding:"required,shortDescription"`
	Price            string `json:"price" binding:"required,currencyString"`
}
