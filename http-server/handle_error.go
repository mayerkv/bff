package http_server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

func handleError(ctx *gin.Context, err error) {
	message := err.Error()
	code := http.StatusInternalServerError

	switch err {
	case io.EOF:
		code = http.StatusBadRequest
	}

	switch err.(type) {
	case validator.ValidationErrors:
		code = http.StatusBadRequest
	}

	ctx.JSON(code, gin.H{
		"error": message,
	})
}
