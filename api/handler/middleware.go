package handler

import (
	"fmt"
	"net/http"

	"github.com/maxmurjon/auth-api/pkg/helper"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		value := c.GetHeader("Authorization")

		info, err := helper.ParseClaims(value, h.cfg.SekretKey)
		if err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		}
		fmt.Println(info)

		c.Set("Auth", info)

		c.Next()
	}
}
