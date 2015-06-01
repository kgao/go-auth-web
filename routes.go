package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
  Route{
    "Hello",
    "GET",
    "/hello",
    Hello,
  },
	Route{
		"Hello",
		"GET",
		"/boom/test",
		BoomTest,
	},
	Route{
		"AuthOnlyGet",
		"GET",
		"/auth/get",
		AuthOnlyGet,
	},
	Route{
		"AuthOnlyPost",
		"POST",
		"/auth/post",
		AuthOnlyPost,
	},
}
