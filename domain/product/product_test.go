package product_test

import (
	"testing"

	"github.com/victoorraphael/tavern/domain/product"
)

func TestProduct_NewProduct(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		description string
		price       float64
		errExpected error
	}

	testCases := []testCase{
		{
			test:        "Empty name",
			name:        "",
			errExpected: product.ErrMissingValues,
		},
		{
			test:        "Create New Product",
			name:        "test",
			description: "test",
			price:       1.0,
			errExpected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := product.NewProduct(tc.name, tc.description, tc.price)
			if err != tc.errExpected {
				t.Errorf("expected %v, got %v", tc.errExpected, err)
			}
		})
	}
}
