package model

type Investor struct {
	ID             int64
	Name           int64
	StocksPosition []*InvestorStocksPosition
}

func NewInvestor(id int64) *Investor {
	return &Investor{
		ID:             id,
		StocksPosition: []*InvestorStocksPosition{},
	}
}

type InvestorStocksPosition struct {
	StockID string
	Shares  int64
}

func NewInvestorStocksPosition(stockID string, shares int64) *InvestorStocksPosition {
	return &InvestorStocksPosition{
		StockID: stockID,
		Shares:  shares,
	}
}

func (i *Investor) AddStockPosition(StockPosition *InvestorStocksPosition) {
	i.StocksPosition = append(i.StocksPosition, StockPosition)
}

func (i *Investor) UpdateStockPosition(stockID string, shares int64) {
	stockPosition := i.GetStockPosition(stockID)

	if stockPosition == nil {
		i.StocksPosition = append(i.StocksPosition, NewInvestorStocksPosition(stockID, shares))
	} else {
		stockPosition.Shares += shares
	}
}

func (i *Investor) GetStockPosition(stockID string) *InvestorStocksPosition {
	for _, stockPosition := range i.StocksPosition {
		if stockPosition.StockID == stockID {
			return stockPosition
		}
	}
	return nil
}
