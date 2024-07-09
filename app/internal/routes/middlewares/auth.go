package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func KataCheckPermCertificate(c *gin.Context) {
	cert := c.Request.Header.Get("Perm-Certificate")
	if cert == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "You don't use certificate or cartificate is invalid",
			"success": false,
		})
		c.Abort()
		return
	}

	c.Next()
}
