package model

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuyStock(t *testing.T) {
	stock1 := NewStock("Stock1", "Stock1", 100)

	investor1 := NewInvestor(1)
	investor2 := NewInvestor(2)

	investorStockPosition := NewInvestorStocksPosition("Stock1", 10)
	investor1.AddStockPosition(investorStockPosition)

	wg := sync.WaitGroup{}
	orderChanIn := make(chan *Order)
	orderChanOut := make(chan *Order)

	offerBook := NewOfferBook(orderChanIn, orderChanOut, &wg)
	go offerBook.Trade()

	wg.Add(1)
	order1 := NewOrder(1, investor1, stock1, 5, 5, "SELL")
	orderChanIn <- order1

	order2 := NewOrder(2, investor2, stock1, 5, 5, "BUY")
	orderChanIn <- order2
	wg.Wait()

	assert := assert.New(t)
	assert.Equal("CLOSED", order1.Status, "Order 1 should be 'CLOSED'")
	assert.Equal(int64(0), order1.PendingShares, "Order 1 should have 0 'PendingShares'")
	assert.Equal("CLOSED", order2.Status, "Order 2 should be 'CLOSED'")
	assert.Equal(int64(0), order2.PendingShares, "Order 2 should hae 0 'PendingShares'")

	assert.Equal(int64(5), investorStockPosition.Shares, "Investor 1 should have 5 shares of 'Stock1'")
	assert.Equal(int64(5), investor2.GetStockPosition("Stock1").Shares, "Investor 2 should have 5 shares of 'Stock1'")
}

func TestBuyStockWithDifferentStocks(t *testing.T) {
	stock1 := NewStock("Stock1", "Stock1", 100)
	stock2 := NewStock("Stock2", "Stock2", 100)

	investor1 := NewInvestor(1)
	investor2 := NewInvestor(2)

	investorStockPosition1 := NewInvestorStocksPosition("Stock1", 10)
	investor1.AddStockPosition(investorStockPosition1)

	investorStockPosition2 := NewInvestorStocksPosition("Stock2", 10)
	investor2.AddStockPosition(investorStockPosition2)

	wg := sync.WaitGroup{}
	orderChanIn := make(chan *Order)
	orderChanOut := make(chan *Order)

	offerBook := NewOfferBook(orderChanIn, orderChanOut, &wg)
	go offerBook.Trade()

	order1 := NewOrder(1, investor1, stock1, 5, 5, "SELL")
	orderChanIn <- order1

	order2 := NewOrder(2, investor2, stock2, 5, 5, "BUY")
	orderChanIn <- order2

	assert := assert.New(t)
	assert.Equal("OPEN", order1.Status, "Order 1 should be 'OPEN'")
	assert.Equal(int64(5), order1.PendingShares, "Order 1 should have 5 'PendingShares'")
	assert.Equal("OPEN", order2.Status, "Order 2 should be 'OPEN'")
	assert.Equal(int64(5), order2.PendingShares, "Order 2 should have 5 'PendingShares'")
}

func TestBuyPartialStock(t *testing.T) {
	stock1 := NewStock("Stock1", "Stock1", 100)

	investor1 := NewInvestor(1)
	investor2 := NewInvestor(2)
	investor3 := NewInvestor(3)

	investorStockPosition1 := NewInvestorStocksPosition("Stock1", 3)
	investor1.AddStockPosition(investorStockPosition1)

	investorStockPosition2 := NewInvestorStocksPosition("Stock1", 5)
	investor3.AddStockPosition(investorStockPosition2)

	wg := sync.WaitGroup{}
	orderChanIn := make(chan *Order)
	orderChanOut := make(chan *Order)

	offerBook := NewOfferBook(orderChanIn, orderChanOut, &wg)
	go offerBook.Trade()

	wg.Add(1)

	order1 := NewOrder(1, investor1, stock1, 3, 5.0, "SELL")
	orderChanIn <- order1

	order2 := NewOrder(2, investor2, stock1, 5, 5.0, "BUY")
	orderChanIn <- order2

	go func() {
		for range orderChanOut {

		}
	}()

	wg.Wait()

	assert := assert.New(t)
	assert.Equal("CLOSED", order1.Status, "Order 1 should be 'CLOSED'")
	assert.Equal(int64(0), order1.PendingShares, "Order 1 should have 0 'PendingShares'")

	assert.Equal("OPEN", order2.Status, "Order 2 should be 'OPEN'")
	assert.Equal(int64(2), order2.PendingShares, "Order 2 should have 2 'PendingShares'")

	assert.Equal(int64(0), investorStockPosition1.Shares, "Investor 1 should have 0 shares of 'Stock1'")
	assert.Equal(int64(3), investor2.GetStockPosition("Stock1").Shares, "Investor 2 should have 3 shares of 'Stock1'")

	wg.Add(1)
	order3 := NewOrder(3, investor3, stock1, 2, 5.0, "SELL")
	orderChanIn <- order3
	wg.Wait()

	assert.Equal("CLOSED", order2.Status, "Order 2 should be 'CLOSED'")
	assert.Equal(int64(0), order2.PendingShares, "Order 2 should have 0 'PendingShares'")

	assert.Equal("CLOSED", order3.Status, "Order 3 should be 'CLOSED'")
	assert.Equal(int64(0), order3.PendingShares, "Order 3 should have 0 'PendingShares'")

	assert.Equal(2, len(offerBook.Transactions), "Should have 2 transactions")
	assert.Equal(15.0, float64(offerBook.Transactions[0].Total), "transaction should have price 15")
	assert.Equal(10.0, float64(offerBook.Transactions[1].Total), "transaction should have price 10")
}

func TestBuyWithDifferentPrice(t *testing.T) {
	stock1 := NewStock("Stock1", "Stock1", 100)

	investor1 := NewInvestor(1)
	investor2 := NewInvestor(2)
	investor3 := NewInvestor(3)

	investorStockPosition1 := NewInvestorStocksPosition("Stock1", 3)
	investor1.AddStockPosition(investorStockPosition1)

	investorStockPosition2 := NewInvestorStocksPosition("Stock1", 5)
	investor3.AddStockPosition(investorStockPosition2)

	wg := sync.WaitGroup{}
	orderChanIn := make(chan *Order)
	orderChanOut := make(chan *Order)

	offerBook := NewOfferBook(orderChanIn, orderChanOut, &wg)
	go offerBook.Trade()

	wg.Add(1)

	order1 := NewOrder(1, investor1, stock1, 3, 4.0, "SELL")
	orderChanIn <- order1

	order2 := NewOrder(2, investor2, stock1, 5, 5.0, "BUY")
	orderChanIn <- order2

	go func() {
		for range orderChanOut {

		}
	}()
	wg.Wait()

	assert := assert.New(t)
	assert.Equal("CLOSED", order1.Status, "Order 1 should be 'CLOSED'")
	assert.Equal(int64(0), order1.PendingShares, "Order 1 should have 0 'PendingShares'")

	assert.Equal("OPEN", order2.Status, "Order 2 should be 'OPEN'")
	assert.Equal(int64(2), order2.PendingShares, "Order 2 should have 2 'PendingShares'")

	assert.Equal(int64(0), investorStockPosition1.Shares, "Investor 1 should have 0 shares of 'Stock1'")
	assert.Equal(int64(3), investor2.GetStockPosition("Stock1").Shares, "Investor 2 should have 3 shares of 'Stock1'")

	wg.Add(1)
	order3 := NewOrder(3, investor3, stock1, 3, 4.5, "SELL")
	orderChanIn <- order3

	wg.Wait()

	assert.Equal("OPEN", order3.Status, "Order 3 should be 'OPEN'")
	assert.Equal(int64(1), order3.PendingShares, "Order 3 should have 1 'PendingShares'")

	assert.Equal("CLOSED", order2.Status, "Order 2 should be 'CLOSED'")
	assert.Equal(int64(0), order2.PendingShares, "Order 2 should have 0 'PendingShares'")
}

func TestNoMatch(t *testing.T) {
	stock1 := NewStock("Stock1", "Stock1", 100)

	investor1 := NewInvestor(1)
	investor2 := NewInvestor(2)

	investorStockPosition1 := NewInvestorStocksPosition("Stock1", 3)
	investor1.AddStockPosition(investorStockPosition1)

	wg := sync.WaitGroup{}
	ordersChanIn := make(chan *Order)
	ordersChanOut := make(chan *Order)

	offerBook := NewOfferBook(ordersChanIn, ordersChanOut, &wg)
	go offerBook.Trade()

	wg.Add(0)

	order := NewOrder(1, investor1, stock1, 3, 6.0, "SELL")
	ordersChanIn <- order

	order2 := NewOrder(2, investor2, stock1, 5, 5.0, "BUY")
	ordersChanIn <- order2

	go func() {
		for range ordersChanOut {

		}
	}()
	wg.Wait()

	assert := assert.New(t)
	assert.Equal("OPEN", order.Status, "Order 1 should be 'OPEN'")
	assert.Equal("OPEN", order2.Status, "Order 2 should be 'OPEN'")
	assert.Equal(int64(3), order.PendingShares, "Order 1 should have 3 'PendingShares'")
	assert.Equal(int64(5), order2.PendingShares, "Order 2 should have 5 'PendingShares'")
}
