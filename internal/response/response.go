package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"my-auction-market-api/internal/errors"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Success: true,
		Data:    data,
	})
}

func SuccessWithStatus(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Data:    data,
	})
}

func Error(c *gin.Context, err error) {
	// 检查是否为应用错误
	if appErr, ok := errors.IsAppError(err); ok {
		c.JSON(appErr.HTTPStatus, Response{
			Success: false,
			Error:   appErr.Message,
		})
		return
	}

	// 默认返回500错误
	c.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Error:   err.Error(),
	})
}

func ErrorWithMessage(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, Response{
		Success: false,
		Error:   message,
	})
}

func BadRequest(c *gin.Context, message string) {
	ErrorWithMessage(c, http.StatusBadRequest, message)
}

func NotFound(c *gin.Context, message string) {
	ErrorWithMessage(c, http.StatusNotFound, message)
}

func Unauthorized(c *gin.Context, message string) {
	ErrorWithMessage(c, http.StatusUnauthorized, message)
}

func Forbidden(c *gin.Context, message string) {
	ErrorWithMessage(c, http.StatusForbidden, message)
}

func Created(c *gin.Context, data interface{}) {
	SuccessWithStatus(c, http.StatusCreated, data)
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}

