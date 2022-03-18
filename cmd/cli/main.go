package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/victoorraphael/tavern/domain/product"
	"github.com/victoorraphael/tavern/services/order"
	tavernservice "github.com/victoorraphael/tavern/services/tavern"
)

func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		panic(err)
	}
	peenuts, err := product.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	products := []product.Product{
		beer, peenuts, wine,
	}
	return products
}

func main() {
	products := productInventory()
	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)

	if err != nil {
		log.Fatal(err)
	}

	tavern, err := tavernservice.NewTavern(
		tavernservice.WithOrderService(os),
	)

	if err != nil {
		log.Fatal(err)
	}

	uid, err := os.AddCustomer("Jordan")

	if err != nil {
		log.Fatal(err)
	}

	pID := []uuid.UUID{
		products[0].GetID(),
	}

	total, err := tavern.Order(uid, pID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("ordered successfully, total =", total)
}
