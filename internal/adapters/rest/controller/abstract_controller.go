package controller

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge-customer-api/internal/errutil"
)

type AbstractController struct{}

type HTTPError struct {
	Message string `json:"error"`
}

func (ctrl *AbstractController) Error(c *gin.Context, err error) {
	log.Println(err)
	var e *errutil.Error
	if errors.As(err, &e) {
		switch e.Type {
		case "BUSINESS":
			c.JSON(422, &HTTPError{Message: e.Message})
		case "INPUT":
			c.JSON(400, &HTTPError{Message: e.Message})
		default:
			c.JSON(500, &HTTPError{Message: "Internal Server Error"})
		}
	} else {
		c.JSON(500, &HTTPError{Message: "Internal Server Error"})
	}
}
