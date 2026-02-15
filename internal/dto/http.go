package dto

import "github.com/gin-gonic/gin"

type BaseResponse struct {
	Status bool   `json:"status"`
	Data   any    `json:"data,omitempty"`
	Err    string `json:"error,omitempty"`
}

func JSON(c *gin.Context, statusCode int, data any, err string) {
	c.JSON(statusCode, BaseResponse{
		Status: err == "",
		Data:   data,
		Err:    err,
	})
}
