package main

import (
	"./router"
	"github.com/gin-gonic/gin"
	// "log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./tmp/*")
	router.Init(r)
	// r.Run(":443")
	s := &http.Server{
		Addr:           ":443",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
