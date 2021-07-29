package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jjonline/study_golang/validator/controller"
	"net/http"
	"time"
)

// controller instance
var (
	testValidator controller.TestValidator
)

func main() {
	var engine = gin.Default()

	// route setting
	engine.GET("/", testValidator.ValidRule)

	server := &http.Server{
		Addr:           ":9080",
		Handler:        engine,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20, //1MB
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
