package server

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

	"example.com/Sinezx/words-server/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func register(c *gin.Context) {
	var request RegReq
	c.BindJSON(&request)
	err := regValidAndFormat(request)
	if err != nil {
		ErrorHandler(c, err)
		return
	}
	_, err = db.QueryUserByAccount(request.Account)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// if account not exist, add current account
		_, err = db.InsertUser(request.Account, request.Password)
		if err != nil {
			ErrorHandler(c, err)
			return
		}
		c.JSON(http.StatusOK, &gin.H{
			"message": "register success",
		})
	} else {
		if err != nil {
			ErrorHandler(c, err)
			return
		}
		c.JSON(http.StatusOK, &gin.H{
			"message": "account already register",
		})
	}
}

func regValidAndFormat(req RegReq) error {
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
