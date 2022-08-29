package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func errorResponse(err error) gin.H {
	return gin.H{"message": err.Error()}
}

func Handler(c *gin.Context, code int, err error) {
	fmt.Printf("error %s, %s", strconv.Itoa(code), err.Error())
	if code == 500 {
		c.JSON(500, map[string]string{
			"message": "internal server error",
		})
	} else {
		c.JSON(code, errorResponse(err))
	}
}
