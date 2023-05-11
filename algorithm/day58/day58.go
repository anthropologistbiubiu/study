package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"github.com/asim/go-micro/v3/util/http"
)

// 匿名结构体的使用
func main() {
	a := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{"sunweiming", 999}
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	http.HandleFunc("/hello", AddUser)
	http.ListenAndServe(":8080", nil)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	//data := ` {"name":"sunweiming","age":999}`

	student := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(student)
	w.Write([]byte("ok"))
}
