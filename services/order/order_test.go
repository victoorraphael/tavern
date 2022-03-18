package order

import (
	"testing"

	"github.com/google/uuid"
	"github.com/victoorraphael/tavern/domain/customer"
	"github.com/victoorraphael/tavern/domain/product"
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

func TestOrder_NewOrder(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	cust, err := customer.NewCustomer("Joel")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	custPr := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetId(), custPr)
	if err != nil {
		t.Error(err)
	}
}
