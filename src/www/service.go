package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayhellName( w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数

	fmt.Println(r.Form)

	fmt.Println("path",r.URL.Path)

	fmt.Println("scheme",r.URL.Scheme)

	fmt.Println(r.Form["url_long"])

	for k,v := range r.Form{
		fmt.Println("key:",k)

		fmt.Println("val:",strings.Join(v,""))
	}

	fmt.Fprintf(w,"Hello astaxie!")

}

func main ()  {
	http.HandleFunc("/",sayhellName)

	err := http.ListenAndServe(":9090",nil)

	if err != nil {
		log.Fatal("ListenAndServer",err)
	}

}

func man() {
	fmt.Println("你好啊")
}
