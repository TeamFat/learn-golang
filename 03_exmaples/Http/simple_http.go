/**
 * datetime: 2017-06-09
 * Go http包 simple web_server
 * https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/03.4.md
 */
package main

import (
	"fmt"
	"net/http"
)

// MyMux struct
type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		hello(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello router")
}

func main() {

	mux := &MyMux{}
	http.ListenAndServe(":8899", mux)

}

//$ go run simple_http.go
