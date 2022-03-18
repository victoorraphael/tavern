//Package memory is a in-memory implementation of the customer repository
package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/victoorraphael/tavern/domain/customer"
)

type MemoryRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

//New is a factory to create a new instance of MemoryRepository
func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

//Implementing the CustomerResposiory Interface
func (m *MemoryRepository) Get(id uuid.UUID) (customer.Customer, error) {
	if customer, ok := m.customers[id]; ok {
		return customer, nil
	}
	return customer.Customer{}, customer.ErrCustomerNotFound
}

func (m *MemoryRepository) Add(c customer.Customer) error {
	if m.customers == nil {
		m.Lock()
		m.customers = make(map[uuid.UUID]customer.Customer)
		m.Unlock()
	}
	//Check if customer already exists in repository
	if _, ok := m.customers[c.GetId()]; ok {
		return customer.ErrFailedToAddCustomer
	}

	m.Lock()
	m.customers[c.GetId()] = c
	m.Unlock()
	return nil
}

func (m *MemoryRepository) Update(c customer.Customer) error {
	if _, ok := m.customers[c.GetId()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrUpdateCustomer)
	}

	m.Lock()
	m.customers[c.GetId()] = c
	m.Unlock()
	return nil
}
