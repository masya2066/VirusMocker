package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"virus_mocker/app/internal/consumer"
	"virus_mocker/app/pkg/files"

	"github.com/gin-gonic/gin"
)

func (a Api) createFilePTMS(c *gin.Context) {
	apiKey := c.GetHeader("X-API-KEY")

	fmt.Print(apiKey)

	fileBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read file",
		})
		return
	}

	file, header, err := files.ByteSliceToMultipartFile(fileBytes, "asdsd")
	if err != nil {
		fmt.Println("Read file error")
		return
	}

	consumer.PtmsChecker(file, header, err)

}
