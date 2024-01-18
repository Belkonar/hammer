package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func globalHandler(w http.ResponseWriter, r *http.Request) {
	uriParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/"), "/")

	prefix := uriParts[0]
	uriParts = uriParts[1:]
	newPath := "/" + strings.Join(uriParts, "/")

	fmt.Println(prefix, newPath)

	client := http.Client{}

	hopByHopHeaders := map[string]bool{
		"Connection":          true,
		"Keep-Alive":          true,
		"Proxy-Authenticate":  true,
		"Proxy-Authorization": true,
		"Te":                  true,
		"Trailers":            true,
		"Transfer-Encoding":   true,
		"Upgrade":             true,
	}

	targetHost := "https://google.com"

	req, _ := http.NewRequest(r.Method, fmt.Sprintf("%s%s", targetHost, newPath), r.Body)

	for key, parts := range r.Header {
		if !hopByHopHeaders[key] {
			for _, val := range parts {
				req.Header.Add(key, val)
			}
		}
	}

	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	for key, parts := range resp.Header {
		if !hopByHopHeaders[key] {
			for _, val := range parts {
				w.Header().Add(key, val)
			}
		}
	}

	w.WriteHeader(resp.StatusCode)

	io.Copy(w, resp.Body)
}

func main() {
	fmt.Println("Hello World!")

	globalHandler := http.HandlerFunc(globalHandler)

	http.ListenAndServe("127.0.0.1:8080", globalHandler)
}
