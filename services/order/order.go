package order

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/victoorraphael/tavern/domain/customer"
	"github.com/victoorraphael/tavern/domain/customer/memory"
	"github.com/victoorraphael/tavern/domain/customer/mongo"
	"github.com/victoorraphael/tavern/domain/product"
	prdmemory "github.com/victoorraphael/tavern/domain/product/memory"
)

type OrderConfiguration func(os *OrderService) error

//A service can hold multiple repositories and other services
type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

func WithCustomerRepository(c customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = c
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	repo := memory.New()
	return WithCustomerRepository(repo)
}

func WithMemoryProductRepository(p []product.Product) OrderConfiguration {
	return func(os *OrderService) error {
		repo := prdmemory.New()
		for _, pr := range p {
			if err := repo.Add(pr); err != nil {
				return err
			}
		}
		os.products = repo
		return nil
	}
}

func WithMongoCustomerRepository(connstring string) OrderConfiguration {
	return func(os *OrderService) error {
		repo, err := mongo.New(context.Background(), connstring)
		if err != nil {
			return err
		}
		os.customers = repo
		return nil
	}
}

func (o *OrderService) CreateOrder(customer uuid.UUID, productsID []uuid.UUID) (float64, error) {
	//Get customer
	c, err := o.customers.Get(customer)
	if err != nil {
		return 0, err
	}
	//Get products
	var products []product.Product
	var total float64

	for _, pID := range productsID {
		pr, err := o.products.GetById(pID)
		if err != nil {
			return 0, err
		}
		products = append(products, pr)
		total += pr.GetPrice()
	}

	log.Printf("Client: %s has ordered %d products", c.GetId(), len(products))

	return total, nil
}

func (o *OrderService) AddCustomer(name string) (uuid.UUID, error) {
	c, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}

	if err := o.customers.Add(c); err != nil {
		return uuid.Nil, err
	}

	return c.GetId(), nil
}
