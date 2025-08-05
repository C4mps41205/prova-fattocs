package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func OK(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{StatusCode: http.StatusOK, Message: message, Data: data})
}

func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{StatusCode: http.StatusCreated, Message: message, Data: data})
}

func BadRequest(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusBadRequest, Response{StatusCode: http.StatusBadRequest, Message: message, Data: data})
}

func NotFound(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusNotFound, Response{StatusCode: http.StatusNotFound, Message: message, Data: data})
}

func InternalServerError(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusInternalServerError, Response{StatusCode: http.StatusInternalServerError, Message: message, Data: data})
}
