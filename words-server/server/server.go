package server

import (
	"net/http"

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

	//datasoure connect
	v1.POST("/dbconnect", dbconnect)
	v1.GET("/disdbconnect", disdbconnect)

	// add handlers
	wordGoup := v1.Group("/word")
	wordGoup.Use(func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("dbconnect") != true {
			c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"message": "please request after datasource connect",
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
