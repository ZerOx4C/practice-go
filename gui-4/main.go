package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"

	"github.com/zserge/lorca"
)

//go:embed res/*
var assets embed.FS

func main() {
	documentRoot, err := fs.Sub(assets, "res")
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.FS(documentRoot)))
	go http.ListenAndServe(":8080", nil)

	ui, _ := lorca.New("http://localhost:8080", "", 640, 480)
	defer ui.Close()

	ui.Bind("notify", func(message string) {
		fmt.Printf("message '%s' received!\n", message)

		switch message {
		case "execute":
			value := ui.Eval("getValueOf('input-box')")
			fmt.Printf("input value is '%s'.\n", value.String())

		default:
			fmt.Println("unknown message.")
		}
	})

	<-ui.Done()
}
