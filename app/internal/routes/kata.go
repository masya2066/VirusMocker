package routes

import (
	"net/http"
	"strings"
	"virus_mocker/app/internal/consumer"
	"virus_mocker/app/internal/db"
	"virus_mocker/app/internal/models/response"

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
		ScanId:   scanId,
		State:    db.KataProcessing,
		SensorId: "",
	}

	if err := a.db.Create(file).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"success": false,
		})
		return
	}

	fl, hd, err := c.Request.FormFile("content")

	go consumer.KataChecker(scanId, fl, hd, err)

	c.JSON(http.StatusOK, gin.H{
		"message": scanId,
		"success": true,
	})
}

func (a Api) GetFiles(c *gin.Context) {
	states := c.Query("state")

	stateList := strings.Split(states, ",")

	var scans []db.KataFile
	if len(stateList) == 0 {
		if err := a.db.Find(&scans).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"success": false,
			})
			return
		}
	} else {
		if err := a.db.Where("state IN ?", stateList).Find(&scans).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
				"success": false,
			})
			return
		}
	}

	var scanIds []db.FileState
	for _, scan := range scans {
		scanIds = append(scanIds, db.FileState{
			ScanId: scan.ScanId,
			State:  scan.State,
		})
	}

	c.JSON(http.StatusOK, response.ScanFilesResult{
		Scans: scanIds,
	})
}

func (a Api) DeleteFile(c *gin.Context) {
	scanId := c.Param("scan_id")
	if scanId == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Undefined scan id",
			"success": false,
		})
		return
	}

	del := a.db.Model(db.KataFile{}).Where("scan_id = ?", scanId).Delete(&db.KataFile{})
	if del.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": del.Error,
			"success": false,
		})
		c.Abort()
		return
	}
	if del.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
			"success": false,
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": scanId,
		"success": true,
	})
}
