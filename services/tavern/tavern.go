package tavern

import (
	"github.com/google/uuid"
	"github.com/victoorraphael/tavern/services/order"
)

type TavernConfiguration func(tv *Tavern) error

type Tavern struct {
	OrderService *order.OrderService
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	tv := &Tavern{}
	for _, cfg := range cfgs {
		if err := cfg(tv); err != nil {
			return nil, err
		}
	}
	return tv, nil
}

func WithOrderService(os *order.OrderService) TavernConfiguration {
	return func(tv *Tavern) error {
		tv.OrderService = os
		return nil
	}
}

func (t *Tavern) Order(customerID uuid.UUID, productsID []uuid.UUID) (float64, error) {
	total, err := t.OrderService.CreateOrder(customerID, productsID)
	if err != nil {
		return 0, err
	}
	return total, nil
}
