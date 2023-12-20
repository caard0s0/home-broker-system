package model

import (
	"container/heap"
	"sync"
)

type OfferBook struct {
	Order         []*Order
	Transactions  []*Transaction
	OrdersChanIn  chan *Order
	OrdersChanOut chan *Order
	wg            *sync.WaitGroup
}

func NewOfferBook(orderChanIn, orderChanOut chan *Order, wg *sync.WaitGroup) *OfferBook {
	return &OfferBook{
		Order:         []*Order{},
		Transactions:  []*Transaction{},
		OrdersChanIn:  orderChanIn,
		OrdersChanOut: orderChanOut,
		wg:            wg,
	}
}

func (ob *OfferBook) Trade() {
	buyOrders := make(map[string]*OrderQueue)
	sellOrders := make(map[string]*OrderQueue)

	for order := range ob.OrdersChanIn {
		stock := order.Stock.ID

		if buyOrders[stock] == nil {
			buyOrders[stock] = NewOrderQueue()
			heap.Init(buyOrders[stock])
		}

		if sellOrders[stock] == nil {
			sellOrders[stock] = NewOrderQueue()
			heap.Init(sellOrders[stock])
		}

		if order.OrderType == "BUY" {
			buyOrders[stock].Push(order)
			if sellOrders[stock].Len() > 0 && sellOrders[stock].Orders[0].Price <= order.Price {
				sellOrder := sellOrders[stock].Pop().(*Order)
				if sellOrder.PendingShares > 0 {
					transaction := NewTransaction(sellOrder, order, order.Shares, sellOrder.Price)
					ob.AddTransaction(transaction, ob.wg)
					sellOrder.Transactions = append(sellOrder.Transactions, transaction)
					order.Transactions = append(order.Transactions, transaction)

					ob.OrdersChanOut <- sellOrder
					ob.OrdersChanOut <- order

					if sellOrder.PendingShares > 0 {
						sellOrders[stock].Push(sellOrder)
					}
				}
			}
		}

		if order.OrderType == "SELL" {
			sellOrders[stock].Push(order)
			if buyOrders[stock].Len() > 0 && buyOrders[stock].Orders[0].Price >= order.Price {
				buyOrder := buyOrders[stock].Pop().(*Order)
				if buyOrder.PendingShares > 0 {
					transaction := NewTransaction(order, buyOrder, order.Shares, buyOrder.Price)
					ob.AddTransaction(transaction, ob.wg)
					buyOrder.Transactions = append(buyOrder.Transactions, transaction)
					order.Transactions = append(order.Transactions, transaction)

					ob.OrdersChanOut <- buyOrder
					ob.OrdersChanOut <- order

					if buyOrder.PendingShares > 0 {
						buyOrders[stock].Push(buyOrder)
					}
				}
			}
		}
	}
}

func (ob *OfferBook) AddTransaction(transaction *Transaction, wg *sync.WaitGroup) {
	defer wg.Done()

	sellingShares := transaction.SellingOrder.PendingShares
	buyingShares := transaction.BuyingOrder.PendingShares

	minShares := sellingShares
	if buyingShares < minShares {
		minShares = buyingShares
	}

	transaction.SellingOrder.Investor.UpdateStockPosition(transaction.SellingOrder.Stock.ID, -minShares)
	transaction.AddSellOrderPendingShares(-minShares)

	transaction.BuyingOrder.Investor.UpdateStockPosition(transaction.BuyingOrder.Stock.ID, minShares)
	transaction.AddBuyOrderPendingShares(-minShares)

	transaction.CalculateTotal(minShares)

	transaction.CloseBuyOrder()
	transaction.CloseSellOrder()

	ob.Transactions = append(ob.Transactions, transaction)
}
