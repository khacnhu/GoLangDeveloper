package middlewares

import (
	"fmt"
	"go-tutorial/internals/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckMiddleware(c *gin.Context) {
	headers := c.GetHeader("Authorization")

	fmt.Println("headers = ", headers)

	if headers == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "credentials in header middleware",
		})
	}

	token := strings.Split(headers, " ")

	if len(token) < 2 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Token not provided or short token",
		})
		return
	}

	data, err := utils.TokenCheck((token[1]))
	fmt.Println("get email in data ", data)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Claims not matched !!!",
		})
		return
	}

	role, roleOk := data["role"].(string)

	if roleOk {
		fmt.Printf("Email: %s\n", role)
	} else {
		fmt.Println("Email not found or not a string")
	}

	// write again to add field role to config author
	if role != "user" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Do you have role to accces in url",
		})
		return
	}

	c.Next()
}
