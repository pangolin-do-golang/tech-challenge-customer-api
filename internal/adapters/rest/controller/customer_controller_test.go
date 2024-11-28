package controller_test

import (
	"errors"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/core/customer"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/adapters/rest/controller"
	"github.com/stretchr/testify/assert"
)

func TestCreateCustomerSuccessfully(t *testing.T) {
	mockService := new(mocks.IService)
	ctrl := controller.NewCustomerController(mockService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/customer", strings.NewReader(`{"name":"John Doe","cpf":"12345678901","email":"john.doe@example.com","age":30}`))
	c.Request.Header.Set("Content-Type", "application/json")

	mockService.On("Create", customer.Customer{
		Name:  "John Doe",
		Cpf:   "12345678901",
		Email: "john.doe@example.com",
		Age:   30,
	}).Return(&customer.Customer{
		Id:    uuid.New(),
		Name:  "John Doe",
		Cpf:   "12345678901",
		Email: "john.doe@example.com",
		Age:   30,
	}, nil)

	ctrl.Create(c)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateCustomerWithInvalidPayload(t *testing.T) {
	ctrl := controller.NewCustomerController(nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/customer", strings.NewReader(`{"name":"J","cpf":"123","email":"invalid","age":17}`))
	c.Request.Header.Set("Content-Type", "application/json")

	ctrl.Create(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestUpdateCustomerSuccessfully(t *testing.T) {
	mockService := new(mocks.IService)
	ctrl := controller.NewCustomerController(mockService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	id := uuid.New()
	c.Params = gin.Params{{Key: "id", Value: id.String()}}
	c.Request = httptest.NewRequest("PUT", "/customer/"+id.String(), strings.NewReader(`{"name":"John Doe","cpf":"12345678901","email":"john.doe@example.com","age":30}`))
	c.Request.Header.Set("Content-Type", "application/json")

	mockService.On("Update", id, customer.Customer{
		Name:  "John Doe",
		Cpf:   "12345678901",
		Email: "john.doe@example.com",
		Age:   30,
	}).Return(&customer.Customer{
		Id:    id,
		Name:  "John Doe",
		Cpf:   "12345678901",
		Email: "john.doe@example.com",
		Age:   30,
	}, nil)

	ctrl.Update(c)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestUpdateCustomerWithInvalidId(t *testing.T) {
	ctrl := controller.NewCustomerController(nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "invalid-uuid"}}
	c.Request = httptest.NewRequest("PUT", "/customer/invalid-uuid", strings.NewReader(`{"name":"John Doe","cpf":"12345678901","email":"john.doe@example.com","age":30}`))
	c.Request.Header.Set("Content-Type", "application/json")

	ctrl.Update(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestDeleteCustomerSuccessfully(t *testing.T) {
	mockService := new(mocks.IService)
	ctrl := controller.NewCustomerController(mockService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	id := uuid.New()
	c.Params = gin.Params{{Key: "id", Value: id.String()}}
	c.Request = httptest.NewRequest("DELETE", "/customer/"+id.String(), nil)

	mockService.On("Delete", id).Return(nil)

	ctrl.Delete(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteCustomerWithInvalidId(t *testing.T) {
	ctrl := controller.NewCustomerController(nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "invalid-uuid"}}
	c.Request = httptest.NewRequest("DELETE", "/customer/invalid-uuid", nil)

	ctrl.Delete(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetAllCustomersSuccessfully(t *testing.T) {
	mockService := new(mocks.IService)
	ctrl := controller.NewCustomerController(mockService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/customer", nil)

	mockService.On("GetAll").Return([]customer.Customer{
		{
			Id:    uuid.New(),
			Name:  "John Doe",
			Cpf:   "12345678901",
			Email: "john.doe@example.com",
			Age:   30,
		},
	}, nil)

	ctrl.GetAll(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCustomerByCpfSuccessfully(t *testing.T) {
	mockService := new(mocks.IService)
	ctrl := controller.NewCustomerController(mockService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "cpf", Value: "12345678901"}}
	c.Request = httptest.NewRequest("GET", "/customer/12345678901", nil)

	mockService.On("GetByCpf", "12345678901").Return(&customer.Customer{
		Id:    uuid.New(),
		Name:  "John Doe",
		Cpf:   "12345678901",
		Email: "john.doe@example.com",
		Age:   30,
	}, nil)

	ctrl.GetByCpf(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCustomerByCpfNotFound(t *testing.T) {
	mockService := new(mocks.IService)
	ctrl := controller.NewCustomerController(mockService)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "cpf", Value: "12345678901"}}
	c.Request = httptest.NewRequest("GET", "/customer/12345678901", nil)

	mockService.On("GetByCpf", "12345678901").Return(nil, errors.New("customer not found"))

	ctrl.GetByCpf(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
