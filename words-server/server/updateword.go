package server

import (
	"net/http"

	"example.com/Sinezx/words-server/db"
	"example.com/Sinezx/words-server/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type UpdateUserWord struct {
	ID     uint `json:"id"`
	Status int  `json:"status"`
}

func updateword(c *gin.Context) {
	session := sessions.Default(c)
	var updateUserWord UpdateUserWord
	c.BindJSON(&updateUserWord)
	switch updateUserWord.Status {
	case util.Remember:
		err := db.UpdateUserWordRate(updateUserWord.ID)
		if err == nil {
			util.InfoFormat("[session:%s]->userword: %d rate update", session.ID(), updateUserWord.ID)
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
