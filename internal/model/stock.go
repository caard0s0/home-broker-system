package model

type Stock struct {
	ID           string
	Name         string
	MarketVolume int64
}

func NewStock(id string, name string, marketVolume int64) *Stock {
	return &Stock{
		ID:           id,
		Name:         name,
		MarketVolume: marketVolume,
	}
}
