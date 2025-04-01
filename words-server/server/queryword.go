package server

import (
	"errors"

	"example.com/Sinezx/words-server/db"
	"github.com/gin-gonic/gin"
)

type QueryReq struct {
	Page     int `json:"page"`
	PageSize int `json:"pagesize"`
}

type QueryWord struct {
	ID          uint    `json:"id"`
	Subject     string  `json:"subject"`
	Translation string  `json:"translation"`
	Rate        float64 `json:"rate"`
}

type QueryResp struct {
	Total int64       `json:"total"`
	Words []QueryWord `json:"words"`
}

func queryword(c *gin.Context) {
	var queryReq QueryReq
	c.BindJSON(&queryReq)
	err := valid(&queryReq)
	if err == nil {
		// query words by limit
		offset := queryReq.PageSize * (queryReq.Page - 1)
		total, res, err := db.Query(offset, queryReq.PageSize)
		if err == nil {
			queryResp := QueryResp{Total: total, Words: swap(res, total)}
			StatusOK(c, &queryResp)
		} else {
			StatusBadRequest(c, &gin.H{"message": err.Error()})
		}
	} else {
		StatusBadRequest(c, &gin.H{"message": err.Error()})
	}
}

func valid(queryReq *QueryReq) error {
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
		queryWords[i].Subject = w.Subject
		queryWords[i].Translation = w.Translation
		queryWords[i].Rate = w.Rate
	}
	return queryWords
}
