package model

const (
	StatusOpen   = "OPEN"
	StatusClosed = "CLOSED"
)

type Order struct {
	ID            int64
	Investor      *Investor
	Stock         *Stock
	Shares        int64
	PendingShares int64
	Price         float64
	OrderType     string
	Status        string
	Transactions  []*Transaction
}

func NewOrder(orderID int64, investor *Investor, stock *Stock, shares int64, price float64, orderType string) *Order {
	return &Order{
		ID:            orderID,
		Investor:      investor,
		Stock:         stock,
		Shares:        shares,
		PendingShares: shares,
		Price:         price,
		OrderType:     orderType,
		Status:        StatusOpen,
		Transactions:  []*Transaction{},
	}
}
