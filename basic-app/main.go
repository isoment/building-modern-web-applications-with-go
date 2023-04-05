package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		n, err := fmt.Fprintf(res, "Hello World")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Bytes written:%d\n", n)
	})

	_ = http.ListenAndServe(":8008", nil)
}
