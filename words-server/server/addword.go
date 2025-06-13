package server

import (
	"errors"
	"net/http"
	"strings"

	"example.com/Sinezx/words-server/db"
	"example.com/Sinezx/words-server/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AddWord struct {
	SourceText string `json:"source_text"`
	TargetText string `json:"target_text"`
}

func addword(c *gin.Context) {
	session := sessions.Default(c)
	var addWord AddWord
	c.BindJSON(&addWord)
	err := addWordValid(addWord)
	if err == nil {
		var word db.Word
		word.UserId = session.Get(util.SessionUserIdKey).(uint)
		word.SourceText = addWord.SourceText
		word.TargetText = addWord.TargetText
		id, err := db.InsertWord(&word)
		if err == nil {
			util.InfoFormat("[session:%s]->word insert success, id: %d", session.ID(), id)
		} else {
			util.InfoFormat("[session:%s]->word insert fail: %s", session.ID(), err.Error())
		}
		if err == nil {
			c.JSON(http.StatusOK, &gin.H{
				"id": id,
			})
		} else {
			ErrorHandler(c, err)
		}
	} else {
		ErrorHandler(c, err)
	}
}

func addWordValid(addWord AddWord) error {
	addWord.SourceText = strings.TrimSpace(addWord.SourceText)
	addWord.TargetText = strings.TrimSpace(addWord.TargetText)
	if addWord.SourceText == "" || addWord.TargetText == "" {
		return errors.New("no params")
	} else {
		return nil
	}
}
