package main

import (
	"fmt"
  "io"
  "log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	 fmt.Fprint(w, "Welcome to the Go world!\n")
}


func Hello(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query()) > 0 {
		fmt.Println("got:", r.URL.Query());
		io.WriteString(w, "Go: Hello, "+r.URL.Query()["user"][0]+"!\n")
	}else{
		io.WriteString(w, "Go: Hello! (You can add your name by ?user=<your name>)\n")
	}

}

//TODO: basic auth
func AuthOnlyGet(w http.ResponseWriter, r *http.Request) {
     if r.Method == "GET" {
					BasicAuth(HandlePost)
          return
      }
      http.Error(w, "get only", http.StatusMethodNotAllowed)
}


func AuthOnlyPost(w http.ResponseWriter, r *http.Request) {
      if r.Method == "POST" {
			  	BasicAuth(HandleGet)
          return
      }
      http.Error(w, "post only", http.StatusMethodNotAllowed)
}


func HandlePost(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    log.Println(r.PostForm)
    io.WriteString(w, "post\n")
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    io.WriteString(w, "Get\n")
}
