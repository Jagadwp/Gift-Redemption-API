package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Meta struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type envelope struct {
	Meta       Meta        `json:"meta"`
	Data       interface{} `json:"data,omitempty"`
	Errors     interface{} `json:"errors,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
	TotalPages  int   `json:"total_pages"`
}

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, envelope{
		Meta: Meta{Code: http.StatusOK, Status: "success", Message: message},
		Data: data,
	})
}

func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, envelope{
		Meta: Meta{Code: http.StatusCreated, Status: "success", Message: message},
		Data: data,
	})
}

func SuccessPaginated(c *gin.Context, message string, data interface{}, pagination *Pagination) {
	c.JSON(http.StatusOK, envelope{
		Meta:       Meta{Code: http.StatusOK, Status: "success", Message: message},
		Data:       data,
		Pagination: pagination,
	})
}

func BadRequest(c *gin.Context, message string, errs interface{}) {
	c.JSON(http.StatusBadRequest, envelope{
		Meta:   Meta{Code: http.StatusBadRequest, Status: "error", Message: message},
		Errors: errs,
	})
}

func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, envelope{
		Meta: Meta{Code: http.StatusUnauthorized, Status: "error", Message: message},
	})
}

func Forbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, envelope{
		Meta: Meta{Code: http.StatusForbidden, Status: "error", Message: message},
	})
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, envelope{
		Meta: Meta{Code: http.StatusNotFound, Status: "error", Message: message},
	})
}

func UnprocessableEntity(c *gin.Context, message string, errs interface{}) {
	c.JSON(http.StatusUnprocessableEntity, envelope{
		Meta:   Meta{Code: http.StatusUnprocessableEntity, Status: "error", Message: message},
		Errors: errs,
	})
}

func InternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, envelope{
		Meta: Meta{Code: http.StatusInternalServerError, Status: "error", Message: message},
	})
}
