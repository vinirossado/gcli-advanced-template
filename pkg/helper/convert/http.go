package convert

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ReadRequestBody(c *gin.Context, requestBody any) {
	err := c.ShouldBindJSON(&requestBody)

	if err != nil {
		panic(err.Error())
	}
}

func ConvertToInt(stringValue string) int {
	valueConv, err := strconv.Atoi(stringValue)

	if err != nil {
		panic(fmt.Sprintf("Error converting parameter to int with error: %s", err))
	}

	return valueConv
}

func ReadBodyParam(c *gin.Context, param string) string {
	return c.Params.ByName("token")
}
