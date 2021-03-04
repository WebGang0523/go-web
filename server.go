package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)


func main() {
	http.HandleFunc("/",say)
	err := http.ListenAndServe(":9091",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func say(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	fmt.Println(r.Form) // 在服务端打印请求参数
	fmt.Println("URL:", r.URL.Path)  // 请求 URL
	fmt.Println("Scheme:", r.URL.Scheme)

	for k, v := range r.Form {
		fmt.Println(k, ":", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "你好，学院君！")  // 发送响应到客户端
}
