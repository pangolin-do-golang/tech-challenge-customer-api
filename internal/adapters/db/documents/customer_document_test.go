package documents_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/adapters/db/documents"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestCustomerDocumentBSONReturnsCorrectMap(t *testing.T) {
	doc := documents.CustomerDocument{
		ID:    uuid.New(),
		Name:  "John Doe",
		Cpf:   "12345678901",
		Email: "john.doe@example.com",
		Age:   30,
	}

	expected := bson.M{
		"name":  "John Doe",
		"cpf":   "12345678901",
		"age":   30,
		"email": "john.doe@example.com",
	}

	result := doc.BSON()

	assert.Equal(t, expected, result)
}

func TestCustomerDocumentBSONIDReturnsCorrectMap(t *testing.T) {
	id := uuid.New()
	doc := documents.CustomerDocument{
		ID: id,
	}

	expected := bson.M{"_id": id}

	result := doc.BSONID()

	assert.Equal(t, expected, result)
}

func TestCustomerDocumentBSONHandlesEmptyFields(t *testing.T) {
	doc := documents.CustomerDocument{
		ID:  uuid.New(),
		Cpf: "12345678901",
	}

	expected := bson.M{
		"name":  "",
		"cpf":   "12345678901",
		"age":   0,
		"email": "",
	}

	result := doc.BSON()

	assert.Equal(t, expected, result)
}
