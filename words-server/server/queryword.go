package server

import (
	"errors"
	"net/http"

	"example.com/Sinezx/words-server/db"
	"example.com/Sinezx/words-server/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type QueryReq struct {
	Page     int `json:"page"`
	PageSize int `json:"pagesize"`
}

type QueryWord struct {
	ID         uint    `json:"id"`
	SourceText string  `json:"source_text"`
	TargetText string  `json:"target_text"`
	Rate       float64 `json:"rate"`
}

type QueryResp struct {
	Total int64       `json:"total"`
	Words []QueryWord `json:"words"`
}

func queryword(c *gin.Context) {
	session := sessions.Default(c)
	var queryReq QueryReq
	c.BindJSON(&queryReq)
	err := queryWordValid(&queryReq)
	if err == nil {
		// query words by limit
		offset := queryReq.PageSize * (queryReq.Page - 1)
		total, res, err := db.QueryWordsByUserId(session.Get(util.SessionUserIdKey).(uint), offset, queryReq.PageSize)
		util.InfoFormat("[session:%s]->query Total: %d", session.ID(), total)
		if err == nil {
			queryResp := QueryResp{Total: total, Words: swap(res, total)}
			c.JSON(http.StatusOK, &queryResp)
		} else {
			ErrorHandler(c, err)
		}
	} else {
		ErrorHandler(c, err)
	}
}

func queryWordValid(queryReq *QueryReq) error {
	if queryReq != nil && queryReq.Page > 0 && queryReq.PageSize > 0 {
		return nil
	} else {
		return errors.New("query param invalid")
	}
}

func swap(source []db.Word, len int64) []QueryWord {
	queryWords := make([]QueryWord, len)
	for i, w := range source {
		queryWords[i].ID = w.ID
		queryWords[i].SourceText = w.SourceText
		queryWords[i].TargetText = w.TargetText
		queryWords[i].Rate = w.Rate
	}
	return queryWords
}
