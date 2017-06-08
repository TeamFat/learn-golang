/**
 * datetime: 2017-06-08
 * Go http包 静态服务器
 * http://localhost:8899/static/avatar.jpg
 */
package main

import (
	"net/http"
)

func main() {
	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)
	http.ListenAndServe(":8899", nil)
}
