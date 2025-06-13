package server

import (
	"net/http"

	"example.com/Sinezx/words-server/db"
	"example.com/Sinezx/words-server/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UpdateWord struct {
	ID     uint `json:"id"`
	Status int  `json:"status"`
}

func updateword(c *gin.Context) {
	session := sessions.Default(c)
	var updateWord UpdateWord
	c.BindJSON(&updateWord)
	switch updateWord.Status {
	case util.Remember:
		err := db.UpdateWordRate(updateWord.ID)
		if err == nil {
			util.InfoFormat("[session:%s]->word: %d rate update", session.ID(), updateWord.ID)
			c.JSON(http.StatusOK, &gin.H{
				"message": "word's rate is updated",
			})
		} else {
			ErrorHandler(c, err)
		}
	default:
		c.JSON(http.StatusBadRequest, &gin.H{
			"message": "status illegal",
		})
	}
}
