package main

import (
	"flag"
	"log"
	"os"
)

var env struct {
	addr string
	infoLog  *log.Logger
	errLog *log.Logger
}

func main() {
	// setup env
	flag.StringVar(&env.addr, "addr", "localhost:8000", "Address of the proxy [default: '127.0.0.1:8000'")
	flag.Parse()

	env.infoLog = log.New(os.Stdout, "[INFO :]  ", log.Ldate|log.Ltime)
	env.errLog = log.New(os.Stderr, "[ERROR:]  ", log.Ldate|log.Ltime|log.Lshortfile)
}
