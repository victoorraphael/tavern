package memory

import (
	"testing"

	"github.com/google/uuid"
	"github.com/victoorraphael/tavern/domain/product"
)

func TestMemoryProductRepository_Add(t *testing.T) {
	repo := New()
	product, err := product.NewProduct("Beer", "Good", 1.99)
	if err != nil {
		t.Error(err)
	}
	err = repo.Add(product)
	if err != nil {
		t.Error(err)
	}
	if len(repo.products) < 1 {
		t.Errorf("expected > 1, got %v", len(repo.products))
	}
}

func TestMemoryProductRepository_Get(t *testing.T) {
	repo := New()
	p, err := product.NewProduct("Beer", "Good", 1.99)
	if err != nil {
		t.Error(err)
	}
	err = repo.Add(p)
	if err != nil {
		t.Error(err)
	}
	if len(repo.products) < 1 {
		t.Errorf("expected >= 1, got %v", len(repo.products))
	}

	type testCase struct {
		test        string
		id          uuid.UUID
		errExpected error
	}

	testCases := []testCase{
		{
			test:        "Get product by ID",
			id:          p.GetID(),
			errExpected: nil,
		},
		{
			test:        "Get non existent product",
			id:          uuid.New(),
			errExpected: product.ErrProductNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := repo.GetById(tc.id)
			if err != tc.errExpected {
				t.Errorf("expected %v, got %v", tc.errExpected, err)
			}
		})
	}
}

func TestMemoryProductRepository_Delete(t *testing.T) {
	repo := New()
	p, err := product.NewProduct("Beer", "Good", 1.99)
	if err != nil {
		t.Error(err)
	}
	err = repo.Add(p)
	if err != nil {
		t.Error(err)
	}
	if len(repo.products) < 1 {
		t.Errorf("expected >= 1, got %v", len(repo.products))
	}

	type testCase struct {
		test        string
		id          uuid.UUID
		errExpected error
	}

	testCases := []testCase{
		{
			test:        "Delete product by ID",
			id:          p.GetID(),
			errExpected: nil,
		},
		{
			test:        "Delete non existent product",
			id:          uuid.New(),
			errExpected: product.ErrProductNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := repo.Delete(tc.id)
			if err != tc.errExpected {
				t.Errorf("expected %v, got %v", tc.errExpected, err)
			}
		})
	}
}
