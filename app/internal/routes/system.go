package routes

import (
	"github.com/gin-gonic/gin"
)

func (a Api) Ping(r *gin.Context) {
	r.JSON(200, gin.H{
		"message": "pong",
	})
}
