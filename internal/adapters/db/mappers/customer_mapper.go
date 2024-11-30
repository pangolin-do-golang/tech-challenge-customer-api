package mappers

import (
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/db/documents"
	"github.com/pangolin-do-golang/tech-challenge/internal/core/customer"
)

func MapCustomerEntityToDocument(entity *customer.Customer) *documents.CustomerDocument {
	document := &documents.CustomerDocument{
		ID:    entity.Id,
		Cpf:   entity.Cpf,
		Name:  entity.Name,
		Email: entity.Email,
		Age:   entity.Age,
	}

	return document
}

func MapCustomerDocumentToEntity(document *documents.CustomerDocument) *customer.Customer {
	return &customer.Customer{
		Id:    document.ID,
		Name:  document.Name,
		Cpf:   document.Cpf,
		Email: document.Email,
		Age:   document.Age,
	}
}

func MapCustomerDocumentToEntityList(documents []*documents.CustomerDocument) []*customer.Customer {
	var customers []*customer.Customer
	for _, document := range documents {
		customers = append(customers, MapCustomerDocumentToEntity(document))
	}
	return customers
}
