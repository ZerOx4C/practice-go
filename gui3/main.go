package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/zserge/lorca"
)

//go:embed res/*
var assets embed.FS

func main() {
	ui, _ := lorca.New("http://localhost:8080", "", 640, 480)
	defer ui.Close()

	documentRoot, err := fs.Sub(assets, "res")
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.FS(documentRoot)))
	go http.ListenAndServe(":8080", nil)

	<-ui.Done()
}
