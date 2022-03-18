package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound      = errors.New("product not found")
	ErrProductAlreadyExists = errors.New("product already exists")
)

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetById(uuid.UUID) (Product, error)
	Add(Product) error
	Update(Product) error
	Delete(uuid.UUID) error
}
