package server

import (
	"net/http"
	"regexp"
	"strconv"

	"example.com/Sinezx/words-server/db"
	"example.com/Sinezx/words-server/util"
	"github.com/gin-gonic/gin"
)

func getwordvoice(c *gin.Context) {
	wordIdStr := c.Param("id")
	matched, err := regexp.Match("\\d+", []byte(wordIdStr))
	if err == nil && matched {
		wordId, _ := strconv.Atoi(wordIdStr)
		if word, err := db.QueryWordById(uint(wordId)); err == nil {
			if word != nil && word.VoicePath != "" {
				util.Info(word.VoicePath)
				c.File(word.VoicePath)
			} else {
				c.JSON(http.StatusOK, &gin.H{
					"message": "no word's voice source",
				})
			}
		} else {
			ErrorHandler(c, err)
		}
	} else {
		ErrorHandler(c, err)
	}
}
