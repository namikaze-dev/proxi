package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var env struct {
	addr    string
	infoLog *log.Logger
	errLog  *log.Logger
}

func main() {
	// setup env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	env.addr = ":" + port

	env.infoLog = log.New(os.Stdout, "[INFO:] ", log.Ldate|log.Ltime)
	env.errLog = log.New(os.Stderr, "[ERROR:] ", log.Ldate|log.Ltime|log.Lshortfile)

	// setup server
	proxyHandler := &proxy{
		client: http.DefaultClient,
	}

	serv := &http.Server{
		Addr:         env.addr,
		Handler:      proxyHandler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     env.errLog,
	}

	env.infoLog.Println("proxy server running on", env.addr)
	err := serv.ListenAndServe()
	if err != nil {
		env.errLog.Fatal(err)
	}
}
