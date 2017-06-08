/**
 * datetime: 2017-06-08
 * Go http包 hello jsser
 */
package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "text/html")
	res.Header().Set("Server", "jsser")
	io.WriteString(res, `<!doctype html>
<html>
<head>
<meta charset="utf-8">
 <title>hello go</title>
</head>
<body>
hello jsser!
</body>
</html>`)
}

func main() {
	//注册路由
	http.HandleFunc("/", hello)
	//监听端口,并且启动服务
	http.ListenAndServe(":8899", nil)

}
