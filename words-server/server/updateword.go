package server

import (
	"example.com/Sinezx/words-server/db"
	"github.com/gin-gonic/gin"
)

const (
	Remember = 0
)

type UpdateWord struct {
	ID     uint `json:"id"`
	Status int  `json:"status"`
}

func updateword(c *gin.Context) {
	var updateWord UpdateWord
	c.BindJSON(&updateWord)
	switch updateWord.Status {
	case Remember:
		word, err := db.UpdateWordRate(updateWord.ID)
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
