package mappers_test

import (
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/core/customer"
	"testing"

	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/adapters/db/documents"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/adapters/db/mappers"
	"github.com/stretchr/testify/assert"
)

func TestMapCustomerEntityToDocumentSuccessfully(t *testing.T) {
	entity := &customer.Customer{
		Id:    uuid.New(),
		Name:  "John Doe",
		Cpf:   "12345678901",
		Email: "john.doe@example.com",
		Age:   30,
	}

	document := mappers.MapCustomerEntityToDocument(entity)

	assert.Equal(t, entity.Id, document.ID)
	assert.Equal(t, entity.Name, document.Name)
	assert.Equal(t, entity.Cpf, document.Cpf)
	assert.Equal(t, entity.Email, document.Email)
	assert.Equal(t, entity.Age, document.Age)
}

func TestMapCustomerDocumentToEntitySuccessfully(t *testing.T) {
	document := &documents.CustomerDocument{
		ID:    uuid.New(),
		Name:  "John Doe",
		Cpf:   "12345678901",
		Email: "john.doe@example.com",
		Age:   30,
	}

	entity := mappers.MapCustomerDocumentToEntity(document)

	assert.Equal(t, document.ID, entity.Id)
	assert.Equal(t, document.Name, entity.Name)
	assert.Equal(t, document.Cpf, entity.Cpf)
	assert.Equal(t, document.Email, entity.Email)
	assert.Equal(t, document.Age, entity.Age)
}

func TestMapCustomerDocumentToEntityListSuccessfully(t *testing.T) {
	documentsList := []*documents.CustomerDocument{
		{
			ID:    uuid.New(),
			Name:  "John Doe",
			Cpf:   "12345678901",
			Email: "john.doe@example.com",
			Age:   30,
		},
		{
			ID:    uuid.New(),
			Name:  "Jane Doe",
			Cpf:   "10987654321",
			Email: "jane.doe@example.com",
			Age:   25,
		},
	}

	entities := mappers.MapCustomerDocumentToEntityList(documentsList)

	assert.Len(t, entities, 2)
	assert.Equal(t, documentsList[0].ID, entities[0].Id)
	assert.Equal(t, documentsList[1].ID, entities[1].Id)
}
