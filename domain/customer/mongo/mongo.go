package mongo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/victoorraphael/tavern/domain/customer"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db       *mongo.Database
	customer *mongo.Collection
}

// mongoCustomer is an internal type that is used to store a CustomerAggregate
// we make an internal struct for this to avoid coupling this mongo implementation to the customeraggregate.
// Mongo uses bson so we add tags for that
type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

//NewFromCustomer takes in a aggregate and converts into a internal structure
func NewFromCustomer(c customer.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetId(),
		Name: c.GetName(),
	}
}

//ToAggregate converts into a customer.Customer
func (m *mongoCustomer) ToAggregate() customer.Customer {
	c := customer.Customer{}
	c.SetId(m.ID)
	c.SetName(m.Name)
	return c
}

//New creates a new mongo repository
func New(ctx context.Context, connstring string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connstring))
	if err != nil {
		return nil, err
	}
	db := client.Database("tavern")
	collection := db.Collection("customers")

	return &MongoRepository{
		db:       db,
		customer: collection,
	}, nil
}

//Implements Customer interface

func (m *MongoRepository) Get(id uuid.UUID) (customer.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), (time.Second * 10))
	defer cancel()

	result := m.customer.FindOne(ctx, bson.M{"id": id})

	var mc mongoCustomer
	err := result.Decode(&mc)
	if err != nil {
		return customer.Customer{}, err
	}

	return mc.ToAggregate(), nil
}

func (m *MongoRepository) Add(c customer.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), (time.Second * 10))
	defer cancel()

	internal := NewFromCustomer(c)
	_, err := m.customer.InsertOne(ctx, internal)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoRepository) Update(customer.Customer) error {
	panic("to implement")
}
