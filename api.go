package main

import (
	"fmt"
	"io"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	 fmt.Fprint(w, "Welcome to the Go world!\n")
}


func Hello(w http.ResponseWriter, r *http.Request) {
   io.WriteString(w, "Go: Hello, World!\n")
}

//TODO: basic auth

func GetWithAuth(w http.ResponseWriter, r *http.Request) {
     if r.Method == "GET" {
          //handler(w, r)
          return
      }
      http.Error(w, "get only", http.StatusMethodNotAllowed)
}


func PostWithAuth(w http.ResponseWriter, r *http.Request) {
      if r.Method == "POST" {
          //handler(w, r)
          return
      }
      http.Error(w, "post only", http.StatusMethodNotAllowed)
}
