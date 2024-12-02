package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/adapters/rest/middleware"
	"github.com/stretchr/testify/assert"
)

func TestOptionsRequestReturns204(t *testing.T) {
	router := gin.Default()
	router.Use(middleware.CorsMiddleware())
	router.GET("/test", func(c *gin.Context) {})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 204, w.Code)
}

func TestGetRequestSetsCorsHeaders(t *testing.T) {
	router := gin.Default()
	router.Use(middleware.CorsMiddleware())
	router.GET("/test", func(c *gin.Context) {})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "GET, POST, PUT, DELETE", w.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Content-Type, Authorization", w.Header().Get("Access-Control-Allow-Headers"))
}

func TestPostRequestSetsCorsHeaders(t *testing.T) {
	router := gin.Default()
	router.Use(middleware.CorsMiddleware())
	router.POST("/test", func(c *gin.Context) {})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "GET, POST, PUT, DELETE", w.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Content-Type, Authorization", w.Header().Get("Access-Control-Allow-Headers"))
}
