package server_test

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/adapters/rest/handler"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/mocks"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/adapters/rest/server"
	"github.com/stretchr/testify/assert"
)

func TestServeStartsServerSuccessfully(t *testing.T) {
	customerService := new(mocks.IService)
	rs := server.NewRestServer(&server.RestServerOptions{
		CustomerService: customerService,
	})

	go func() {
		rs.Serve()
	}()
}

func TestServeHandlesCorsMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	service := new(mocks.IService)
	service.On("GetAll", mock.Anything).Return(nil, nil)
	handler.RegisterCustomerHandlers(router, service)

	req, _ := http.NewRequest(http.MethodGet, "/customer", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
