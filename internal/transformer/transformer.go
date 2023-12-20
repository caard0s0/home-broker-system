package transformer

import (
	"github.com/caard0s0/home-broker-system/internal/dto"
	"github.com/caard0s0/home-broker-system/internal/model"
)

func TransformInput(input dto.TradeInput) *model.Order {
	stock := model.NewStock(input.StockID, input.StockID, 1000)
	investor := model.NewInvestor(input.InvestorID)
	order := model.NewOrder(input.OrderID, investor, stock, input.Shares, input.Price, input.OrderType)
	if input.CurrentShares > 0 {
		stockPosition := model.NewInvestorStocksPosition(input.StockID, input.CurrentShares)
		investor.AddStockPosition(stockPosition)
	}
	return order
}

func TransformOutput(order *model.Order) *dto.OrderOutput {
	output := &dto.OrderOutput{
		OrderID:    order.ID,
		InvestorID: order.Investor.ID,
		StockID:    order.Stock.ID,
		OrderType:  order.OrderType,
		Status:     order.Status,
		Partial:    order.PendingShares,
		Shares:     order.Shares,
	}

	var transactionsOutput []*dto.TransactionOutput
	for _, t := range order.Transactions {
		transactionOutput := &dto.TransactionOutput{
			TransactionID: t.ID,
			BuyerID:       t.BuyingOrder.Investor.ID,
			SellerID:      t.SellingOrder.Investor.ID,
			StockID:       t.SellingOrder.Stock.ID,
			Price:         t.Price,
			Shares:        t.BuyingOrder.Shares - t.BuyingOrder.PendingShares,
		}
		transactionsOutput = append(transactionsOutput, transactionOutput)
	}
	output.TransactionsOutput = transactionsOutput
	return output
}
