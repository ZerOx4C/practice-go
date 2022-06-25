package main

import (
	_ "embed"
	"net/url"

	"github.com/zserge/lorca"
)

//go:embed index.html
var html string

func main() {
	ui, _ := lorca.New("data:text/html,"+url.PathEscape(html), "", 640, 480)
	defer ui.Close()

	<-ui.Done()
}
