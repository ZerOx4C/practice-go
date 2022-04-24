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

	mainUi, _ := lorca.New("http://localhost:8080/main.html", "", 640, 480)
	defer mainUi.Close()

	subUis := []lorca.UI{}
	defer func() {
		for _, subUi := range subUis {
			subUi.Close()
		}
	}()

	mainUi.Bind("notify", func() {
		subUi, _ := lorca.New("http://localhost:8080/sub.html", "", 320, 240)
		subUi.Bind("notify", func() {
			fmt.Println("action!")
		})
		subUis = append(subUis, subUi)
	})

	<-mainUi.Done()
}
