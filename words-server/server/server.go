package server

import (
	"net/http"

	"example.com/Sinezx/words-server/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func Helloworld(c *gin.Context) {
	c.JSON(http.StatusOK, "helloworld")
}

func Run() {
	r := gin.Default()

	// example
	v1 := r.Group("/api/v1")
	exampleGroup := v1.Group("/example")
	exampleGroup.GET("/helloworld", Helloworld)

	// session controller
	store := memstore.NewStore([]byte("store"))
	v1.Use(sessions.Sessions("mystore", store))

	v1.POST("/sayhi", sayhi)

	// add handlers
	wordGoup := v1.Group("/word")
	wordGoup.Use(func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get(util.SessionUserIdKey) == nil {
			c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"message": "Ladies and Gentlemen, say hi to me",
			})
			c.Abort()
		} else {
			c.Next()
		}
	})
	wordGoup.POST("/queryword", queryword)
	wordGoup.POST("/addword", addword)
	wordGoup.POST("/updateword", updateword)

	r.Run()
}

func StatusOK(c *gin.Context, resp ...any) {
	c.JSON(http.StatusOK, resp)
}

func StatusBadRequest(c *gin.Context, resp ...any) {
	c.JSON(http.StatusBadRequest, resp)
}

func ErrorHandler(c *gin.Context, err error) {
	StatusBadRequest(c, &gin.H{
		"message": err.Error(),
	})
}
