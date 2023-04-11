package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		value := c.GetHeader("Password")
		fmt.Println(value)
		if value != h.cfg.SecretKey {
			c.AbortWithError(http.StatusForbidden, errors.New("password invalid"))
			return
		}

		c.Next()
	}
}
	