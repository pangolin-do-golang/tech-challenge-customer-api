package controller_test

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/adapters/rest/controller"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/errutil"
	"github.com/stretchr/testify/assert"
)

func TestBusinessErrorReturns422(t *testing.T) {
	ctrl := &controller.AbstractController{}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	err := errutil.NewBusinessError(errors.New("business error"), "business error occurred")

	ctrl.Error(c, err)

	assert.Equal(t, 422, w.Code)
	assert.JSONEq(t, `{"error":"business error occurred"}`, w.Body.String())
}

func TestInputErrorReturns400(t *testing.T) {
	ctrl := &controller.AbstractController{}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	err := errutil.NewInputError(errors.New("Bad Request"))

	ctrl.Error(c, err)

	assert.Equal(t, 400, w.Code)
	assert.JSONEq(t, `{"error":"Bad Request"}`, w.Body.String())
}

func TestUnknownErrorTypeReturns500(t *testing.T) {
	ctrl := &controller.AbstractController{}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	err := &errutil.Error{Message: "unknown error", Type: "UNKNOWN"}

	ctrl.Error(c, err)

	assert.Equal(t, 500, w.Code)
	assert.JSONEq(t, `{"error":"Internal Server Error"}`, w.Body.String())
}

func TestNonCustomErrorReturns500(t *testing.T) {
	ctrl := &controller.AbstractController{}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	err := errors.New("standard error")

	ctrl.Error(c, err)

	assert.Equal(t, 500, w.Code)
	assert.JSONEq(t, `{"error":"Internal Server Error"}`, w.Body.String())
}
