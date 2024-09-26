package routes

import (
	"net/http"
	"virus_mocker/app/internal/models/response"
	"virus_mocker/app/pkg/generator"

	"github.com/gin-gonic/gin"
)

func (a Api) CreateFileMS(c *gin.Context) {
	c.JSON(http.StatusOK, response.CreateFileMS{
		Errors: nil,
		Data: response.CreateFileMSData{
			FileUri: "https://s3.amazonaws.com/virus-mocker/0f6f3e8f-0c2b-4e9b-8f6c-1b5c6b8d9e10",
			Ttl:     3600,
		},
	})
}

func (a Api) CreateScanTaskMS(c *gin.Context) {

	uuid, err := generator.Uuid()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	c.JSON(http.StatusOK, response.CreateScanTaskMS{
		Errors: nil,
		Data: response.CreateScanTaskMSData{
			ScanId: uuid,
		},
	})
}
