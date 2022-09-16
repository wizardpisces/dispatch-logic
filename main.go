package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	dispatch "github.com/wizardpisces/dispatch-logic/text"

	"github.com/wizardpisces/dispatch-logic/structures"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认是不会解析的
	// fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	// fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])

	fmt.Println("listening ", dispatch.Portal{}.Text(structures.Normal, "name"))

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, fmt.Sprintf("Hello %s, %d!", structures.Normal, structures.Normal)) // 这个写入到 w 的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName) // 设置访问的路由
	fmt.Println("server is run on http://localhost:9090")
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
