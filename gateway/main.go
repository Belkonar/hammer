package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func globalHandler(w http.ResponseWriter, r *http.Request) {
	uriParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/"), "/")

	prefix := uriParts[0]
	uriParts = uriParts[1:]
	newPath := "/" + strings.Join(uriParts, "/")

	fmt.Println(prefix, newPath) // just so it won't complain about unused variables

	proxy := httputil.ReverseProxy{ // TODO: Make a factory so this can be reused
		Rewrite: func(r *httputil.ProxyRequest) {
			url, err := url.Parse("https://google.com" + newPath)
			r.SetXForwarded()

			if err != nil {
				panic(err) // FIXME: handle error
			}

			r.Out.URL = url
			r.Out.Host = url.Host // Super annoying but entirely necessary
		},
	}

	proxy.ServeHTTP(w, r)
}

func main() {
	fmt.Println("Hello World!")

	globalHandler := http.HandlerFunc(globalHandler)

	http.ListenAndServe("127.0.0.1:8080", globalHandler)
}
