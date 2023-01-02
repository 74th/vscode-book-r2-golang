package main

import (
	"flag"
	"os"

	"github.com/74th/vscode-book-r2-golang/server"
)

func main() {
	var (
		webroot string
		addr    string
	)

	flag.StringVar(&webroot, "webroot", "./public", "web root path")
	flag.StringVar(&addr, "addr", "0.0.0.0:8000", "server addr")
	flag.Parse()

	svr := server.New(addr, webroot)
	err := svr.Serve()
	if err != nil {
		os.Exit(1)
	}
}
