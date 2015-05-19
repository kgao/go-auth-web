package main

import (
	"log"
	"io"
	"net/http"
)

func up(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Server is up.")
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:    ":9090",
		Handler: &myHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
  //public:
	mux["/"] = up
  //private:
  mux["/hello"] = HandleIndex
	//TODO: BUGFIX - To call a func from other file even under same package,user go run *.go : https://github.com/go-lang-plugin-org/go-lang-idea-plugin/issues/555

	mux["/post"] = PostOnly(BasicAuth(HandlePost))

	mux["/post"] = GetOnly(BasicAuth(HandleJSON))

	log.Println("Server is up on localhost:9090")

	server.ListenAndServe()

}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My server: "+r.URL.String())
}
