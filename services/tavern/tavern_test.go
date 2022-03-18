package tavern

import (
	"testing"

	"github.com/google/uuid"
	"github.com/victoorraphael/tavern/domain/product"
	"github.com/victoorraphael/tavern/services/order"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}
	peenuts, err := product.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}
	products := []product.Product{
		beer, peenuts, wine,
	}
	return products
}

func TestTavern(t *testing.T) {
	products := init_products(t)
	os, err := order.NewOrderService(
		//if i wish to use mongo instead memory
		//WithMongoCustomerRepository("connectionstring")
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	uid, err := os.AddCustomer("Eder")
	if err != nil {
		t.Error(err)
	}

	productsID := []uuid.UUID{
		products[0].GetID(),
		products[1].GetID(),
	}

	total, err := tavern.Order(uid, productsID)
	if err != nil {
		t.Error(err)
	}

	if total == 0 {
		t.Errorf("expect != 0, got %v", total)
	}
}
