package main

import (
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
)

func main() {
	doc, err := htmlquery.LoadURL("https://zerox4c.github.io/practice-go/table.html")
	if err != nil {
		panic(err)
	}

	table := [][]string{}
	for _, rowNode := range htmlquery.Find(doc, "//*[@id='list']/tbody/tr") {
		row := []string{}
		for _, cellNode := range htmlquery.Find(rowNode, "td") {
			row = append(row, htmlquery.InnerText(cellNode))
		}
		table = append(table, row)
	}

	for _, row := range table {
		fmt.Println(strings.Join(row, "\t"))
	}
}
