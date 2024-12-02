package repositories_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/adapters/db/documents"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/mocks"
	"go.mongodb.org/mongo-driver/v2/bson"
	"testing"

	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/adapters/db/repositories"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/core/customer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func TestCreateCustomerSuccessfully(t *testing.T) {
	mockCollection := new(mocks.ICollection)
	repo := repositories.NewMongoCustomerRepository(mockCollection)
	cust := &customer.Customer{Name: "John Doe", Cpf: "12345678901", Email: "john.doe@example.com", Age: 30}

	mockCollection.On("InsertOne", mock.Anything, mock.Anything).Return(&mongo.InsertOneResult{}, nil)

	result, err := repo.Create(context.Background(), cust)

	assert.NoError(t, err)
	assert.Equal(t, cust.Name, result.Name)
}

func TestCreateCustomerWithInvalidData(t *testing.T) {
	mockCollection := new(mocks.ICollection)
	repo := repositories.NewMongoCustomerRepository(mockCollection)
	cust := &customer.Customer{Name: "", Cpf: "12345678901", Email: "john.doe@example.com", Age: 30}

	mockCollection.On("InsertOne", mock.Anything, mock.Anything).Return(nil, errors.New("invalid data"))

	result, err := repo.Create(context.Background(), cust)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUpdateCustomerNotFound(t *testing.T) {
	mockCollection := new(mocks.ICollection)
	repo := repositories.NewMongoCustomerRepository(mockCollection)
	cust := &customer.Customer{Id: uuid.New(), Name: "John Doe", Cpf: "12345678901", Email: "john.doe@example.com", Age: 30}
	document := &documents.CustomerDocument{ID: cust.Id, Name: "John Doe", Cpf: "12345678901", Email: "john.doe@example.com", Age: 30}

	mockCollection.On("FindOneAndUpdate", mock.Anything, document.BSONID(), mock.Anything, mock.Anything).Return(&mongo.SingleResult{})

	result, err := repo.Update(context.Background(), cust)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestDeleteCustomerSuccessfully(t *testing.T) {
	mockCollection := new(mocks.ICollection)
	repo := repositories.NewMongoCustomerRepository(mockCollection)
	id := uuid.New()

	mockCollection.On("DeleteOne", mock.Anything, bson.M{"_id": id}, mock.Anything).Return(&mongo.DeleteResult{}, nil)

	err := repo.Delete(context.Background(), id)

	assert.NoError(t, err)
}

func TestDeleteCustomerNotFound(t *testing.T) {
	mockCollection := new(mocks.ICollection)
	repo := repositories.NewMongoCustomerRepository(mockCollection)
	id := uuid.New()

	mockCollection.On("DeleteOne", mock.Anything, bson.M{"_id": id}, mock.Anything).Return(nil, mongo.ErrNoDocuments)

	err := repo.Delete(context.Background(), id)

	assert.Error(t, err)
}

func TestGetCustomerByCpfNotFound(t *testing.T) {
	mockCollection := new(mocks.ICollection)
	repo := repositories.NewMongoCustomerRepository(mockCollection)

	mockCollection.On("FindOne", mock.Anything, bson.M{"cpf": "12345678901"}, mock.Anything).Return(&mongo.SingleResult{})

	result, err := repo.GetByCpf(context.Background(), "12345678901")

	assert.Error(t, err)
	assert.Nil(t, result)
}
