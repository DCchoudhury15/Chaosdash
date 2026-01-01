package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, err := url.Parse("https://www.google.com")
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.Host = target.Host
	}

	fmt.Println("Starting Chaos Proxy on port 8080...")
	if err := http.ListenAndServe(":8080", proxy); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
