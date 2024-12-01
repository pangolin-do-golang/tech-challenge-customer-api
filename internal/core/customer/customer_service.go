package customer

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type Service struct {
	repository IRepository
}

func NewService(customerRepository IRepository) IService {
	return &Service{
		repository: customerRepository,
	}
}

func (s *Service) Create(ctx context.Context, customer *Customer) (*Customer, error) {
	existingCustomer, _ := s.GetByCpf(ctx, customer.Cpf)

	if existingCustomer != nil {
		return nil, errors.New("entered cpf is already registered in our system")
	}

	return s.repository.Create(ctx, customer)
}

func (s *Service) Update(ctx context.Context, customer *Customer) (*Customer, error) {
	return s.repository.Update(ctx, customer)
}

func (s *Service) Delete(ctx context.Context, customerId uuid.UUID) error {
	return s.repository.Delete(ctx, customerId)
}

func (s *Service) GetAll(ctx context.Context) ([]*Customer, error) {
	return s.repository.GetAll(ctx)
}

func (s *Service) GetByCpf(ctx context.Context, cpf string) (*Customer, error) {
	return s.repository.GetByCpf(ctx, cpf)
}
