package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/jchv/go-webview2"
)

//go:embed embed/*
var asset embed.FS

func main() {
	docRoot, err := fs.Sub(asset, "embed")

	server := new(Server)
	err = server.Start(http.FileServer(http.FS(docRoot)))
	if err != nil {
		panic(err)
	}

	webview := webview2.New(false)
	defer webview.Destroy()

	webview.Navigate(server.Url)
	webview.Run()

	server.Close()
}
