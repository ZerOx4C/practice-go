package main

import (
	"embed"
	"encoding/csv"
	"io/fs"
	"net/http"
	"os"
	"strconv"

	"github.com/zserge/lorca"
)

//go:embed res/*
var assets embed.FS
var ui lorca.UI

func main() {
	ui, _ = lorca.New("http://localhost:8080", "", 640, 480)
	defer ui.Close()

	ui.Bind("save", save)

	documentRoot, err := fs.Sub(assets, "res")
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.FS(documentRoot)))
	go http.ListenAndServe(":8080", nil)

	<-ui.Done()
}

func save() {
	tableValue := ui.Eval("getList()").Array()
	table := [][]string{
		{"id", "name", "age"},
	}

	for _, rowValue := range tableValue {
		row := []string{
			strconv.Itoa(rowValue.Object()["id"].Int()),
			rowValue.Object()["name"].String(),
			strconv.Itoa(rowValue.Object()["age"].Int()),
		}
		table = append(table, row)
	}

	file, err := os.Create("output.csv")
	if err != nil {
		panic(err)
	}

	csvWriter := csv.NewWriter(file)
	csvWriter.WriteAll(table)
	file.Close()

	ui.Eval("alert('保存したよ')")
}
