package server

import (
	"net/http"
	"path/filepath"

	"example.com/Sinezx/words-server/util"
	"github.com/gin-gonic/gin"
)

func uploadwordvoice(c *gin.Context) {
	file, _ := c.FormFile("file")
	c.SaveUploadedFile(file, util.Config.VoiceFolder+filepath.Base(file.Filename))
	c.JSON(http.StatusOK, &gin.H{
		"message": "file upload done",
	})
}
