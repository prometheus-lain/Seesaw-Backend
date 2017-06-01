package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func handle(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Server", "Seesaw Backend")
	uri := request.URL.Path[1:]
	if uri != "" {
		//not homepage
		var page string
		if strings.Contains(uri, "/") {
			//subpage
			uri_parts := strings.Split(uri, "/")
			page = uri_parts[0]
		} else {
			page = uri
		}
		switch page {
		case "view":
			fmt.Fprintf(w, "aa\n")
		case "login":
			fmt.Fprintf(w, "please log in: \n")
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "404\n")
		}
	} else {
		fmt.Fprintf(w, "Hello, welcome\n")
	}
	fmt.Fprintf(w, "<br>"+uri+"\n")
}

func main() {
	http.HandleFunc("/", handle)
	//certificate generated with generate_cert.go, change with letsencrypt soon(tm)
	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	log.Panic(err)
}
