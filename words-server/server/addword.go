package server

import (
	"time"

	"example.com/Sinezx/words-server/db"
	"example.com/Sinezx/words-server/util"
	"github.com/gin-gonic/gin"
)

type AddWord struct {
	Subject     string `json:"subject"`
	Translation string `json:"translation"`
}

func addword(c *gin.Context) {
	var addWord AddWord
	c.BindJSON(&addWord)
	var word db.Word
	word.Subject = addWord.Subject
	word.Translation = addWord.Translation
	word.RateUpAt = time.Now().UTC()
	word.Rate = util.TheEbbinghausForgettingCurve(float64(1) / 60)
	id, err := db.Insert(&word)
	if err == nil {
		StatusOK(c, &gin.H{
			"id": id,
		})
	} else {
		StatusBadRequest(c, &gin.H{
			"message": err.Error(),
		})
	}
}
