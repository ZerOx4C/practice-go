package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/zserge/lorca"
)

//go:embed res/*
var assets embed.FS
var ui lorca.UI

func main() {
	documentRoot, err := fs.Sub(assets, "res")
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.FS(documentRoot)))
	go http.ListenAndServe(":8080", nil)

	ui, _ = lorca.New("http://localhost:8080", "", 640, 480)
	defer ui.Close()

	ui.Bind("fetchString", fetchString)
	ui.Bind("fetchInt", fetchInt)
	ui.Bind("fetchArray", fetchArray)
	ui.Bind("fetchTable", fetchTable)
	ui.Bind("fetchComplex", fetchComplex)

	<-ui.Done()
}

func fetchString() (string, error) {
	return "hello", nil
}

func fetchInt() (int, error) {
	return 1228, nil
}

func fetchArray() ([]string, error) {
	return []string{"foo", "bar", "baz"}, nil
}

func fetchTable() (map[string]int, error) {
	return map[string]int{"foo": 123, "bar": 456, "baz": 789}, nil
}

func fetchComplex() (map[string]interface{}, error) {
	return map[string]interface{}{
		"foo": 123,
		"bar": "hello",
		"baz": map[string]int{
			"hoge": 456,
			"piyo": 789,
		},
	}, nil
}
