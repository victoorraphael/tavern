package customer

//Aggregates package holds aggregates that combines many entities into a full object

import (
	"errors"

	"github.com/google/uuid"
	"github.com/victoorraphael/tavern"
)

//An aggregate not allow access to inner properties
//thats why all the fields are unexported

var (
	ErrInvalidPerson = errors.New("a customer has to have a valid person")
)

//Customer is a aggregate that combines all entities needed to represent a customer
type Customer struct {
	//person is the root entity of a customer
	person       *tavern.Person
	products     []*tavern.Item
	transactions []tavern.Transaction
}

//NewCustomer is a factory to create a New Customer aggregate
//It will validate that name is not empty
func NewCustomer(name string) (Customer, error) {
	// Validate that the name is not empty
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	//Create a new person and generate ID
	person := &tavern.Person{
		Name: name,
		ID:   uuid.New(),
	}

	//Create a customer object
	return Customer{
		person:       person,
		products:     make([]*tavern.Item, 0),
		transactions: make([]tavern.Transaction, 0),
	}, nil
}

func (c *Customer) GetId() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetId(id uuid.UUID) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.ID = id
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.Name = name
}
