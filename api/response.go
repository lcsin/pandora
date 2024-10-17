package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lcsin/pandora/api/message"
)

// Response 响应体封装
type Response struct {
	Code    message.ErrCode `json:"code"`
	Message string          `json:"message"`
	Data    any             `json:"data"`
}

// ResponseOK 请求成功
//
//	@param c
//	@param data
func ResponseOK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code:    0,
		Message: "ok",
		Data:    data,
	})
}

// ResponseError 请求失败，错误码
//
//	@param c
//	@param code
func ResponseError(c *gin.Context, code message.ErrCode) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: code.String(),
	})
}

// ResponseErrorMessage 请求失败，错误信息
//
//	@param c
//	@param code
//	@param msg
func ResponseErrorMessage(c *gin.Context, code message.ErrCode, msg string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
	})
}
