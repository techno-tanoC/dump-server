package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var bs, err = httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
