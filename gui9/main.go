package main

import (
	"fmt"
	"net/http"

	"github.com/jchv/go-webview2"
)

func main() {
	count := 0

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf(`
		<html>
		<body>
		%d<input type="button" value="add" onclick="add()" />
		</body>
		</html>
		`, count)))
	}))
	go http.ListenAndServe(":8080", mux)

	debug := true
	w := webview2.New(debug)
	defer w.Destroy()
	w.Navigate("http://localhost:8080")
	w.Bind("add", func() {
		count++
		w.Eval("location.reload()")
	})
	w.Run()
}
