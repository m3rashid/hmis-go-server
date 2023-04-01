package middlewares

import (
	"net/http"
	"strings"

	"github.com/m3rashid-org/hmis-go-server/utils"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bear := c.Request.Header.Get("Authorization")
		token := strings.Replace(bear, "Bearer ", "", 1)
		sub, err := utils.Decoder(token)
		if err != nil {
			c.Abort()
			c.String(http.StatusUnauthorized, err.Error())
		} else {
			c.Set("sub", sub)
			c.Next()
		}
	}
}
