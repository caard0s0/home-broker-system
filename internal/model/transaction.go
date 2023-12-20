package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID           string
	SellingOrder *Order
	BuyingOrder  *Order
	Shares       int64
	Price        float64
	Total        float64
	CreatedAt    time.Time
}

func NewTransaction(sellingOrder, buyingOrder *Order, shares int64, price float64) *Transaction {
	total := float64(shares) * price
	return &Transaction{
		ID:           uuid.New().String(),
		SellingOrder: sellingOrder,
		BuyingOrder:  buyingOrder,
		Shares:       shares,
		Price:        price,
		Total:        total,
		CreatedAt:    time.Now(),
	}
}

func (t *Transaction) CalculateTotal(shares int64) {
	t.Total = float64(shares) * t.Price
}

func (t *Transaction) CloseBuyOrder() {
	if t.BuyingOrder.PendingShares == 0 {
		t.BuyingOrder.Status = "CLOSED"
	}
}

func (t *Transaction) CloseSellOrder() {
	if t.SellingOrder.PendingShares == 0 {
		t.SellingOrder.Status = "CLOSED"
	}
}

func (t *Transaction) AddBuyOrderPendingShares(shares int64) {
	t.BuyingOrder.PendingShares += shares
}

func (t *Transaction) AddSellOrderPendingShares(shares int64) {
	t.SellingOrder.PendingShares += shares
}
