package server

import (
	"errors"
	"regexp"
	"strings"

	"example.com/Sinezx/words-server/db"
	"example.com/Sinezx/words-server/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SayHiReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func sayhi(c *gin.Context) {
	var request SayHiReq
	c.BindJSON(&request)
	err := sayhiValidAndFormat(request)
	if err != nil {
		ErrorHandler(c, err)
		return
	}
	user, err := db.QueryUserByAccount(request.Account)
	if err != nil {
		if err.Error() == "record not found" {
			// insert user
			userId, err := db.InsertUser(request.Account, request.Password)
			if err != nil {
				ErrorHandler(c, err)
				return
			}
			sessionSaveUserInfo(c, userId)
			StatusOK(c, &gin.H{
				"user_id": userId,
			})
			return
		} else {
			ErrorHandler(c, err)
			return
		}
	}
	//compare password md5
	if user.Password == util.Md5(request.Password) {
		sessionSaveUserInfo(c, user.ID)
		StatusOK(c, &gin.H{
			"user_id": user.ID,
		})
	} else {
		StatusOK(c, &gin.H{
			"message": "password incorrect or not your account",
		})
	}

}

func sayhiValidAndFormat(req SayHiReq) error {
	// format
	req.Account = strings.TrimSpace(req.Account)
	req.Password = strings.TrimSpace(req.Password)
	// valid
	if req.Account == "" || req.Password == "" {
		return errors.New("no params")
	}
	matched, err := regexp.Match("", []byte(req.Account))
	if err != nil || !matched {
		return err
	}
	return nil
}

func sessionSaveUserInfo(c *gin.Context, userId uint) {
	session := sessions.Default(c)
	session.Set(util.SessionUserIdKey, userId)
	session.Save()
}
