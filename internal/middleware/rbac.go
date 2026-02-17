package middleware

import (
	"github.com/gift-redemption/internal/model"
	"github.com/gift-redemption/internal/pkg/response"
	"github.com/gin-gonic/gin"
)

// RequireRole allows access only to users with one of the given roles.
func RequireRole(roles ...model.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get(ContextRole)
		if !exists {
			response.Forbidden(c, "access denied")
			c.Abort()
			return
		}

		for _, r := range roles {
			if model.UserRole(role.(string)) == r {
				c.Next()
				return
			}
		}

		response.Forbidden(c, "you do not have permission to access this resource")
		c.Abort()
	}
}
