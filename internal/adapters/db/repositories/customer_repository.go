package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/adapters/db/documents"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/adapters/db/mappers"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/core/customer"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoCustomerRepository struct {
	collection ICollection
}

func NewMongoCustomerRepository(db ICollection) customer.IRepository {
	return &MongoCustomerRepository{collection: db}
}

func (r *MongoCustomerRepository) Create(ctx context.Context, cust *customer.Customer) (*customer.Customer, error) {
	newID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	document := mappers.MapCustomerEntityToDocument(cust)
	document.ID = newID

	_, err = r.collection.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}

	cust.Id = newID

	return cust, nil
}

func (r *MongoCustomerRepository) Update(ctx context.Context, cust *customer.Customer) (*customer.Customer, error) {
	document := mappers.MapCustomerEntityToDocument(cust)
	update := bson.M{"$set": document.BSON()}

	result := r.collection.FindOneAndUpdate(ctx, document.BSONID(), update, options.FindOneAndUpdate().SetUpsert(false))
	if result.Err() != nil {
		return nil, fmt.Errorf("error updating client: %w", result.Err())
	}

	return cust, nil
}

func (r *MongoCustomerRepository) Delete(ctx context.Context, customerId uuid.UUID) error {
	document := &documents.CustomerDocument{ID: customerId}
	_, err := r.collection.DeleteOne(ctx, document.BSONID())
	if err != nil {
		return err
	}

	return nil
}

func (r *MongoCustomerRepository) GetAll(ctx context.Context) ([]*customer.Customer, error) {
	var customers []*documents.CustomerDocument
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error finding customers: %w", err)
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &customers)
	if err != nil {
		return nil, fmt.Errorf("error decoding customers: %w", err)
	}

	return mappers.MapCustomerDocumentToEntityList(customers), nil
}

func (r *MongoCustomerRepository) GetByCpf(ctx context.Context, customerCpf string) (*customer.Customer, error) {
	var document documents.CustomerDocument
	filter := bson.M{"cpf": customerCpf}

	err := r.collection.FindOne(ctx, filter).Decode(&document)
	if err != nil {
		return nil, err
	}

	return &customer.Customer{
		Id:    document.ID,
		Name:  document.Name,
		Cpf:   document.Cpf,
		Email: document.Email,
		Age:   document.Age,
	}, nil
}

type ICollection interface {
	InsertOne(ctx context.Context, document interface{},
		opts ...options.Lister[options.InsertOneOptions]) (*mongo.InsertOneResult, error)
	DeleteOne(
		ctx context.Context,
		filter interface{},
		opts ...options.Lister[options.DeleteOptions],
	) (*mongo.DeleteResult, error)
	Find(ctx context.Context, filter interface{},
		opts ...options.Lister[options.FindOptions]) (*mongo.Cursor, error)
	FindOne(ctx context.Context, filter interface{},
		opts ...options.Lister[options.FindOneOptions]) *mongo.SingleResult
	FindOneAndUpdate(
		ctx context.Context,
		filter interface{},
		update interface{},
		opts ...options.Lister[options.FindOneAndUpdateOptions]) *mongo.SingleResult
}
