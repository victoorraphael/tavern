package memory

import (
	"sync"

	"github.com/google/uuid"
	"github.com/victoorraphael/tavern/domain/product"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]product.Product),
	}
}

func (m *MemoryProductRepository) GetAll() ([]product.Product, error) {
	var products []product.Product
	for _, p := range m.products {
		products = append(products, p)
	}
	return products, nil
}

func (m *MemoryProductRepository) GetById(id uuid.UUID) (product.Product, error) {
	if _, ok := m.products[id]; !ok {
		return product.Product{}, product.ErrProductNotFound
	}

	return m.products[id], nil
}

func (m *MemoryProductRepository) Add(p product.Product) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[p.GetID()]; ok {
		return product.ErrProductAlreadyExists
	}

	m.products[p.GetID()] = p
	return nil
}

func (m *MemoryProductRepository) Update(p product.Product) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[p.GetID()]; !ok {
		return product.ErrProductNotFound
	}

	m.products[p.GetID()] = p
	return nil
}

func (m *MemoryProductRepository) Delete(id uuid.UUID) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.products[id]; !ok {
		return product.ErrProductNotFound
	}

	delete(m.products, id)
	return nil
}
