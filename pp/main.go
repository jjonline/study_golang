package main

import (
	"net/http"
	"time"
	_ "net/http/pprof"
)

func main()  {
	srv := &http.Server{
		Addr:           ":9080",
		Handler:        nil,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   65 * time.Second,
		MaxHeaderBytes: 1 << 20, //1MB
	}

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello world!"))
	})

	_ = srv.ListenAndServe()
}
