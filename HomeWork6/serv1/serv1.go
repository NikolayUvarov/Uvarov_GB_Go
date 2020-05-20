package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	keys, ok := req.URL.Query()["name"]
	name := "anonymous"
	if ok && len(keys[0]) > 1 {
		name = keys[0]
	}
	io.WriteString(res,
		`<doctype html>
<html>
    <head>
        <title>Hello `+name+`!</title>
    </head>
    <body>
        Hello `+name+`!
    </body>
</html>`)
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hello)

	http.ListenAndServe(":80", nil)
}
