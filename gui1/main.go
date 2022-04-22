package main

import (
	"net/url"

	"github.com/zserge/lorca"
)

func main() {
	ui, _ := lorca.New("data:text/html,"+url.PathEscape(`
	<html>
		<head>
			<title>GUIのサンプルだよ</title>
			<meta charset="UTF-8" />
		</head>
		<body>
			<h1>静的UIだよ</h1>
			<input type="text" placeholder="入力ボックスだよ" />
			<input type="button" value="ボタンだよ" onclick="javascript:alert('押したよ')" />
		</body>
	</html>
	`), "", 640, 480)
	defer ui.Close()

	<-ui.Done()
}
