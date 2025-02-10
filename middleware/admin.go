package middleware

import (
	"server/model/appTypes"
	"server/model/response"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID := utils.GetRoleID(c)

		if roleID != appTypes.Admin {
			response.Forbidden("Access denied. Adimin privileges are required", c)
			c.Abort()
			return
		}

		c.Next()
	}
}
