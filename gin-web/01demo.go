package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		file, _ := ioutil.ReadFile("./hello.txt")
		fmt.Fprintln(w,string(file))
	})

	err := http.ListenAndServe("localhost:9090",nil)
	if err != nil {
		fmt.Printf("http serve failed,err:%v\n",err)
		return
	}
}
