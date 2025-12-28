package page

import (
	"github.com/gin-gonic/gin"
)

type PageQuery struct {
	Page     int `json:"page" form:"page" binding:"omitempty,min=1"`
	PageSize int `json:"pageSize" form:"pageSize" binding:"omitempty,min=1,max=100"`
}

func (q *PageQuery) Bind(c *gin.Context) error {
	if err := c.ShouldBindQuery(q); err != nil {
		c.Error(err)
		return err
	}

	if q.Page == 0 {
		q.Page = 1
	}
	if q.PageSize == 0 {
		q.PageSize = 10
	}

	return nil
}

func (q *PageQuery) Offset() int {
	return (q.Page - 1) * q.PageSize
}

func (q *PageQuery) Limit() int {
	return q.PageSize
}

type PageData struct {
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
	Total    int64       `json:"total"`
	Data     interface{} `json:"data"`
}

func NewPageData(page, pageSize int, total int64, data interface{}) *PageData {
	return &PageData{
		Page:     page,
		PageSize: pageSize,
		Total:    total,
		Data:     data,
	}
}

