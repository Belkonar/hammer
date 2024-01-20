package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func globalHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	proxy := makeProxy("http://example.com")

	proxy.ServeHTTP(w, r)
}

func makeProxy(target string) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{ // TODO: Make a factory so this can be reused
		Rewrite: makeRewriter(target),
	}
}

func makeRewriter(target string) func(*httputil.ProxyRequest) {
	return func(r *httputil.ProxyRequest) {
		target, err := url.Parse(target)

		if err != nil {
			panic(err) // FIXME: handle error
		}

		r.SetURL(target)

		r.SetXForwarded()
		r.Out.Host = target.Host // Super annoying but entirely necessary
		fmt.Println(r.Out.URL)
	}
}

func main() {
	fmt.Println("Hello World!")

	globalHandler := http.HandlerFunc(globalHandler)

	http.ListenAndServe("127.0.0.1:8000", globalHandler)
}
