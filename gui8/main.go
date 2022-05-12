package main

import (
	"net/url"

	"github.com/jchv/go-webview2"
	"github.com/ncruces/zenity"
)

func main() {
	debug := true
	w := webview2.New(debug)
	defer w.Destroy()
	w.SetTitle("Minimal webview example")
	w.SetSize(300, 300, webview2.HintNone)
	w.SetSize(200, 200, webview2.HintMin)
	w.SetSize(400, 400, webview2.HintMax)
	w.Navigate("data:text/html," + url.PathEscape(`
	<html>
		<head>
			<title>GUIのサンプルだよ</title>
			<meta charset="UTF-8" />
		</head>
		<body>
			<h1>静的UIだよ</h1>
			<input type="text" placeholder="入力ボックスだよ" />
			<input type="button" value="ボタンだよ" onclick="javascript:alert('押したよ')" />
			<input type="button" value="ボタンだよ" onclick="javascript:hoge()" />
		</body>
	</html>
	`))
	w.Bind("hoge", func() {
		zenity.SelectFile()
		panic("yay!")
	})
	w.Run()
	println("wow!")
}
