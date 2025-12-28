package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	govalidator "github.com/go-playground/validator/v10"

	appValidator "my-auction-market-api/internal/validator"
)

func ValidationErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				if validationErr, ok := err.Err.(govalidator.ValidationErrors); ok {
					errors := appValidator.FormatValidationErrors(validationErr)
					c.JSON(http.StatusBadRequest, gin.H{
						"success": false,
						"error":   "validation failed",
						"errors":  errors,
					})
					c.Abort()
					return
				}
			}
		}
	}
}

