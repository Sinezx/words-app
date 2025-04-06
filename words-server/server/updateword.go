package server

import (
	"example.com/Sinezx/words-server/db"
	"example.com/Sinezx/words-server/util"
	"github.com/gin-gonic/gin"
)

type UpdateWord struct {
	ID     uint `json:"id"`
	Status int  `json:"status"`
}

func updateword(c *gin.Context) {
	session := util.GetSession(c)
	var updateWord UpdateWord
	c.BindJSON(&updateWord)
	switch updateWord.Status {
	case util.Remember:
		word, err := db.UpdateWordRate(session.DB(), updateWord.ID)
		util.InfoFormat("[session:%s]->word: %d rate update", session.ID(), word.ID)
		if err == nil {
			StatusOK(c, &word)
		} else {
			StatusBadRequest(c, &gin.H{
				"message": err.Error(),
			})
		}
	default:
		StatusBadRequest(c, &gin.H{
			"message": "status illegal",
		})
	}
}
