package customer_test

import (
	"testing"

	"github.com/victoorraphael/tavern/domain/customer"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		errExpected error
	}

	//test cases
	testCases := []testCase{
		{
			test:        "Empty name",
			name:        "",
			errExpected: customer.ErrInvalidPerson,
		},
		{
			test:        "Valid name",
			name:        "Valid name",
			errExpected: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			//Create a new customer
			_, err := customer.NewCustomer(tc.name)
			if err != tc.errExpected {
				t.Errorf("Expected error %v, got %v", tc.errExpected, err)
			}
		})
	}
}
