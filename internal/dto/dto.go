package dto

type TradeInput struct {
	OrderID       int64   `json:"order_id"`
	InvestorID    int64   `json:"investor_id"`
	StockID       string  `json:"stock_id"`
	CurrentShares int64   `json:"current_shares"`
	Shares        int64   `json:"shares"`
	Price         float64 `json:"price"`
	OrderType     string  `json:"order_type"`
}

type OrderOutput struct {
	OrderID            int64                `json:"order_id"`
	InvestorID         int64                `json:"investor_id"`
	StockID            string               `json:"stock_id"`
	OrderType          string               `json:"order_type"`
	Status             string               `json:"status"`
	Partial            int64                `json:"partial"`
	Shares             int64                `json:"shares"`
	TransactionsOutput []*TransactionOutput `json:"transactions"`
}

type TransactionOutput struct {
	TransactionID string  `json:"transaction_id"`
	BuyerID       int64   `json:"buyer_id"`
	SellerID      int64   `json:"seller_id"`
	StockID       string  `json:"stock_id"`
	Price         float64 `json:"price"`
	Shares        int64   `json:"shares"`
}
