package server

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"

	"example.com/Sinezx/words-server/db"
	"example.com/Sinezx/words-server/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AddWord struct {
	SourceText string `form:"source_text"`
	TargetText string `form:"target_text"`
}

func addword(c *gin.Context) {
	session := sessions.Default(c)
	var addWord AddWord
	// c.ShouldBind(&addWord)
	addWord.SourceText = c.Request.FormValue("source_text")
	addWord.TargetText = c.Request.FormValue("target_text")
	err := addWordValid(addWord)
	if err == nil {
		//insert userword
		var userWord db.UserWord
		userWord.UserId = session.Get(util.SessionUserIdKey).(uint)
		userWord.WordId = saveWordVoicAndGetWordId(addWord)

		uwid, err := db.InsertUserWord(&userWord)
		if err == nil {
			util.InfoFormat("[session:%s]->word insert success, id: %d", session.ID(), uwid)
		} else {
			util.InfoFormat("[session:%s]->word insert fail: %s", session.ID(), err.Error())
		}
		if err == nil {
			c.JSON(http.StatusOK, &gin.H{
				"uwid": uwid,
			})
		} else {
			ErrorHandler(c, err)
		}
	} else {
		ErrorHandler(c, err)
	}
}

func saveWordVoicAndGetWordId(addWord AddWord) uint {
	//check the word is exists or not
	wordId := db.QueryWordId(addWord.SourceText)
	if wordId > 0 {
		return wordId
	} else {
		//insert word
		var word db.Word
		word.SourceText = addWord.SourceText
		word.TargetText = addWord.TargetText

		// save binary file to local and record local path
		word.VoicePath = savewordvoice(addWord.SourceText)
		db.InsertWord(&word)

		return word.ID
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

func savewordvoice(sourceText string) string {
	resp, e := http.Get(util.Config.VoiceSourceUrl + sourceText)
	if e != nil {
		util.Info(e.Error())
	} else {
		body, e := io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode > 299 {
			util.InfoFormat("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
		}
		if e != nil {
			util.Info(e.Error())
		}
		localPath := util.Config.VoiceFolder + sourceText
		e = os.WriteFile(localPath, body, 0777)
		if e != nil {
			util.Info(e.Error())
		} else {
			return localPath
		}
	}
	return ""
}
