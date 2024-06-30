package routes

import (
	"net/http"
	"virus_mocker/app/internal/db"

	"github.com/gin-gonic/gin"
)

func (a Api) CreateFile(c *gin.Context) {
	objectType := c.Request.FormValue("objectType")
	if objectType != "file" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Undefined object type",
			"success": false,
		})
		return
	}
	scanId := c.Request.FormValue("scanId")

	file := &db.KataFile{
		ScanId: scanId,
		Status: "processing",
	}

	if err := a.db.Create(file).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": scanId,
		"success": true,
	})
}
