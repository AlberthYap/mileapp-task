package utils

import (
	"task-api/types"

	"github.com/gin-gonic/gin"
)

// Success - send success response
func Success(c *gin.Context, statusCode int, message string, data interface{}) {
  c.JSON(statusCode, types.SuccessResponse{
    Status:  "success",
    Message: message,
    Data:    data,
  })
}

// Fail - send fail response
func Fail(c *gin.Context, statusCode int, message string, data interface{}) {
  c.JSON(statusCode, types.FailResponse{
    Status:  "fail",
    Message: message,
    Data:    data,
  })
}

// Error - send error response
func Error(c *gin.Context, statusCode int, message string, code int, data interface{}) {
  response := types.ErrorResponse{
    Status:  "error",
    Message: message,
  }
  
  if code > 0 {
    response.Code = code
  }
  
  if data != nil {
    response.Data = data
  }
  
  c.JSON(statusCode, response)
}
