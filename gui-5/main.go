package main

import (
	"embed"
	"encoding/csv"
	"io/fs"
	"net/http"
	"os"
	"strconv"

	"github.com/ncruces/zenity"
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

	ui.Bind("save", save)

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

	path, err := openSaveCsvDialog()
	if err != nil {
		return
	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	csvWriter := csv.NewWriter(file)
	csvWriter.WriteAll(table)
	file.Close()

	ui.Eval("alert('保存したよ')")
}

func openSaveCsvDialog() (string, error) {
	options := []zenity.Option{}
	options = append(options, zenity.FileFilter{Name: "CSVファイル (*.csv)", Patterns: []string{"*.csv"}})
	options = append(options, zenity.FileFilter{Name: "すべてのファイル (*.*)", Patterns: []string{"*"}})
	options = append(options, zenity.ConfirmOverwrite())
	return zenity.SelectFileSave(options...)
}
