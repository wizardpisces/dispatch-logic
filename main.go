package main

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	dispatch "github.com/wizardpisces/dispatch-logic/text"

	"github.com/gin-gonic/gin"
	"github.com/wizardpisces/dispatch-logic/structures"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认是不会解析的
	// fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	// fmt.Println("path", r.URL.Path)
	// fmt.Println("scheme", r.URL.Scheme)
	// fmt.Println(r.Form["url_long"])a

	fmt.Println("listening ", dispatch.Portal{}.Text(structures.Normal, "name"))

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, fmt.Sprintf("Hello %s, %d!", structures.Normal, structures.Normal)) // 这个写入到 w 的是输出到客户端的
}

// func main() {
// 	http.HandleFunc("/", sayhelloName) // 设置访问的路由
// 	fmt.Println("server is run on http://localhost:9090")
// 	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }

func defaultRouter(ctx *gin.Context) {
	fmt.Println("listening default", dispatch.Portal{}.Text(structures.Normal, "name"))
	ctx.JSON(http.StatusOK, gin.H{
		"default": fmt.Sprintf("Hello %s, %d!", structures.Normal, structures.Normal),
	})
}

func TestGuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}

func main() {
	// TestGuessingGame()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", defaultRouter)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
