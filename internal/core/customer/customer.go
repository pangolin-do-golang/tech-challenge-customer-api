package customer

import (
	"context"

	"github.com/google/uuid"
)

type Customer struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Cpf   string    `json:"cpf"`
	Email string    `json:"email"`
	Age   int       `json:"age"`
}

type IService interface {
	Create(ctx context.Context, customer *Customer) (*Customer, error)
	Update(ctx context.Context, customer *Customer) (*Customer, error)
	Delete(ctx context.Context, customerId uuid.UUID) error
	GetAll(ctx context.Context) ([]*Customer, error)
	GetByCpf(ctx context.Context, customerCpf string) (*Customer, error)
}

type IRepository interface {
	Create(ctx context.Context, customer *Customer) (*Customer, error)
	Update(ctx context.Context, customer *Customer) (*Customer, error)
	Delete(ctx context.Context, customerId uuid.UUID) error
	GetAll(ctx context.Context) ([]*Customer, error)
	GetByCpf(ctx context.Context, customerCpf string) (*Customer, error)
}
