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

func Run() error {
	r := gin.Default()

	// example
	v1 := r.Group("/api/v1")
	exampleGroup := v1.Group("/example")
	exampleGroup.GET("/helloworld", Helloworld)

	// session controller
	store := memstore.NewStore([]byte("store"))
	v1.Use(sessions.Sessions("mystore", store))

	v1.POST("/reg", register)
	v1.POST("/sayhi", sayhi)

	// add handlers
	wordGroup := v1.Group("/word")
	wordGroup.Use(func(c *gin.Context) {
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
	wordGroup.POST("/queryword", queryword)
	wordGroup.POST("/addword", addword)
	wordGroup.POST("/updateword", updateword)
	wordGroup.POST("/uploadwordvoice", uploadwordvoice)
	wordGroup.GET("/getwordvoice/:id", getwordvoice)

	return r.Run()
}

func ErrorHandler(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, &gin.H{
		"message": err.Error(),
	})
}
