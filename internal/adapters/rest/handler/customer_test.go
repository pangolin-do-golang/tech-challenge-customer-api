package handler_test

import (
	"github.com/pangolin-do-golang/tech-challenge-customer-api/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/adapters/rest/handler"
	"github.com/stretchr/testify/assert"
)

func TestRegisterCustomerHandlersSuccessfully(t *testing.T) {
	mockService := new(mocks.IService)
	mockService.On("GetAll").Return(nil, nil)
	router := gin.Default()
	handler.RegisterCustomerHandlers(router, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/customer", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRegisterCustomerHandlersWithInvalidRoute(t *testing.T) {
	mockService := new(mocks.IService)
	router := gin.Default()
	handler.RegisterCustomerHandlers(router, mockService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/invalid", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
