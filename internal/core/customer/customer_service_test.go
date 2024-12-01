package customer_test

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/core/customer"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/mocks"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func TestService_Create(t *testing.T) {
	c := customer.Customer{
		Name: "Teste",
		Cpf:  "123",
	}

	type fields struct {
		genRepository func() customer.IRepository
	}
	type args struct {
		customer customer.Customer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *customer.Customer
		wantErr bool
	}{
		{
			name: "returns error for existing customer",
			fields: fields{
				genRepository: func() customer.IRepository {
					m := new(mocks.IRepository)
					m.On("GetByCpf", mock.Anything, c.Cpf).Return(nil, errors.New("error"))
					return m
				},
			},
			args: args{
				customer: c,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns error for non-null customer",
			fields: fields{
				genRepository: func() customer.IRepository {
					m := new(mocks.IRepository)
					m.On("GetByCpf", mock.Anything, c.Cpf).Return(&customer.Customer{}, nil)
					return m
				},
			},
			args: args{
				customer: c,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "returns error from customer creation",
			fields: fields{
				genRepository: func() customer.IRepository {
					m := new(mocks.IRepository)
					m.On("GetByCpf", mock.Anything, c.Cpf).Return(nil, nil)
					m.On("Create", mock.Anything, &c).Return(nil, errors.New("error"))
					return m
				},
			},
			args: args{
				customer: c,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "creates customer",
			fields: fields{
				genRepository: func() customer.IRepository {
					m := new(mocks.IRepository)
					m.On("GetByCpf", mock.Anything, c.Cpf).Return(nil, nil)
					m.On("Create", mock.Anything, &c).Return(&c, nil)
					return m
				},
			},
			args: args{
				customer: c,
			},
			want:    &c,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := customer.NewService(tt.fields.genRepository())
			got, err := s.Create(context.TODO(), &tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Update(t *testing.T) {
	c := customer.Customer{
		Id:   uuid.MustParse("cd661160-4624-49dd-a531-ba48dfce45c2"),
		Name: "Ronaldinho",
	}

	updatedCustomer := customer.Customer{
		Id:   uuid.MustParse("cd661160-4624-49dd-a531-ba48dfce45c2"),
		Name: "Ronaldo",
	}

	type fields struct {
		genRepository func() customer.IRepository
	}
	type args struct {
		customerId uuid.UUID
		customer   customer.Customer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *customer.Customer
		wantErr bool
	}{
		{
			name: "returns error from update method",
			fields: fields{
				genRepository: func() customer.IRepository {
					m := new(mocks.IRepository)
					m.On("Update", mock.Anything, &c).Return(nil, errors.New("error"))
					return m
				},
			},
			args: args{
				customer: c,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "updates customer",
			fields: fields{
				genRepository: func() customer.IRepository {
					m := new(mocks.IRepository)
					m.On("Update", mock.Anything, &c).Return(&updatedCustomer, nil)
					return m
				},
			},
			args: args{
				customerId: c.Id,
				customer:   c,
			},
			want:    &updatedCustomer,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := customer.NewService(tt.fields.genRepository())
			got, err := s.Update(context.TODO(), &tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_Delete(t *testing.T) {
	id := uuid.MustParse("d8cdc2b2-3014-4053-941f-2af4fd036a50")
	type fields struct {
		genRepository func() customer.IRepository
	}
	type args struct {
		customerId uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "returns error from delete method",
			fields: fields{
				genRepository: func() customer.IRepository {
					m := new(mocks.IRepository)
					m.On("Delete", mock.Anything, id).Return(errors.New("error"))
					return m
				},
			},
			args: args{
				customerId: id,
			},
			wantErr: true,
		},
		{
			name: "deletes customer",
			fields: fields{
				genRepository: func() customer.IRepository {
					m := new(mocks.IRepository)
					m.On("Delete", mock.Anything, id).Return(nil)
					return m
				},
			},
			args: args{
				customerId: id,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := customer.NewService(tt.fields.genRepository())
			if err := s.Delete(context.TODO(), tt.args.customerId); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_GetAll(t *testing.T) {
	type fields struct {
		genRepository func() customer.IRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*customer.Customer
		wantErr bool
	}{
		{
			name: "returns error from get all method",
			fields: fields{
				genRepository: func() customer.IRepository {
					m := new(mocks.IRepository)
					m.On("GetAll", mock.Anything).Return([]*customer.Customer{}, errors.New("error"))
					return m
				},
			},
			want:    []*customer.Customer{},
			wantErr: true,
		},
		{
			name: "gets all customers",
			fields: fields{
				genRepository: func() customer.IRepository {
					m := new(mocks.IRepository)
					m.On("GetAll", mock.Anything).Return([]*customer.Customer{
						{
							Name: "Ronaldo",
						},
					}, nil)
					return m
				},
			},
			want: []*customer.Customer{
				{
					Name: "Ronaldo",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := customer.NewService(tt.fields.genRepository())
			got, err := s.GetAll(context.TODO())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetByCpf(t *testing.T) {
	cpf := "132"
	type fields struct {
		genRepository func() customer.IRepository
	}
	type args struct {
		cpf string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *customer.Customer
		wantErr bool
	}{
		{
			name: "returns error from get by cpf method",
			fields: fields{
				genRepository: func() customer.IRepository {
					m := new(mocks.IRepository)
					m.On("GetByCpf", mock.Anything, cpf).Return(nil, errors.New("error"))
					return m
				},
			},
			args: args{
				cpf: cpf,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "gets customer by cpf",
			fields: fields{
				genRepository: func() customer.IRepository {
					m := new(mocks.IRepository)
					m.On("GetByCpf", mock.Anything, cpf).Return(&customer.Customer{
						Name: "Ronaldo",
					}, nil)
					return m
				},
			},
			args: args{
				cpf: cpf,
			},
			want:    &customer.Customer{Name: "Ronaldo"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := customer.NewService(tt.fields.genRepository())
			got, err := s.GetByCpf(context.TODO(), tt.args.cpf)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByCpf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByCpf() got = %v, want %v", got, tt.want)
			}
		})
	}
}
