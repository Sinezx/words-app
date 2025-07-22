package server

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func uploadwordvoice(c *gin.Context) {
	file, _ := c.FormFile("file")
	c.SaveUploadedFile(file, filepath.Base(file.Filename))
	c.JSON(http.StatusOK, &gin.H{
		"message": "file upload done",
	})
}
