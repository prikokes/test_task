package models

type MerchListResponse struct {
	Items []Merch `json:"items"`
}

type TransactionsResponse struct {
	Received []TransactionInfo `json:"received"`
	Sent     []TransactionInfo `json:"sent"`
}

type TransactionInfo struct {
	Username string `json:"username"`
	Money    int64  `json:"money"`
}
