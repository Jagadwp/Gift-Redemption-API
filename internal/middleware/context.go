package middleware

import "github.com/gin-gonic/gin"

func GetUserID(c *gin.Context) uint {
	val, _ := c.Get(ContextUserID)
	id, _ := val.(uint)
	return id
}

func GetRole(c *gin.Context) string {
	val, _ := c.Get(ContextRole)
	role, _ := val.(string)
	return role
}
