package main

import (
	"flag"
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
	flag.StringVar(&env.addr, "addr", "localhost:8000", "Address of the proxy [default: '127.0.0.1:8000'")
	flag.Parse()

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
		ErrorLog: env.errLog,
	}

	env.infoLog.Println("proxy server running on", env.addr)
	err := serv.ListenAndServe()
	if err != nil {
		env.errLog.Fatal(err)
	}
}
