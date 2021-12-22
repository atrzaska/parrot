package main

import (
	"github.com/gin-gonic/gin"
  "net/http"
)

func main() {
  gin.SetMode(gin.ReleaseMode)
  body := "<img src='https://i.giphy.com/5KfqXl7ZCKqgE.gif' />"
	router := gin.Default()
	router.GET("/*all", func(c *gin.Context) {
    c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(body))
  })
  router.Run(":4300")
}
