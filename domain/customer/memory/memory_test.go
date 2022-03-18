package memory

import (
	"testing"

	"github.com/google/uuid"
	"github.com/victoorraphael/tavern/domain/customer"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		test      string
		id        uuid.UUID
		expectErr error
	}

	c, err := customer.NewCustomer("Raphael")
	if err != nil {
		t.Fatal(err)
	}
	id := c.GetId()
	mem := MemoryRepository{customers: map[uuid.UUID]customer.Customer{
		id: c,
	}}

	testCases := []testCase{
		{
			test:      "No customer by id",
			id:        uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectErr: customer.ErrCustomerNotFound,
		},
		{
			test:      "Customer by id",
			id:        id,
			expectErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := mem.Get(tc.id)
			if err != tc.expectErr {
				t.Errorf("expected %v, got %v", tc.expectErr, err)
			}
		})
	}
}

func TestMemory_AddCustomer(t *testing.T) {
	type testCase struct {
		test      string
		cust      string
		expectErr error
	}

	testCases := []testCase{
		{
			test:      "Add a customer",
			cust:      "Peter",
			expectErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			repo := MemoryRepository{
				customers: make(map[uuid.UUID]customer.Customer),
			}

			c, err := customer.NewCustomer(tc.cust)
			if err != nil {
				t.Fatal(err)
			}

			err = repo.Add(c)
			if err != tc.expectErr {
				t.Errorf("expected %v, got %v", tc.expectErr, err)
			}

			cust, err := repo.Get(c.GetId())
			if err != nil {
				t.Fatal(err)
			}
			if cust.GetId() != c.GetId() {
				t.Errorf("expected %v, got %v", c.GetId(), cust.GetId())
			}
		})
	}
}
