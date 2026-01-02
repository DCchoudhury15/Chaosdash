package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
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

func ChaosMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if rand.Intn(100) < 50 {
			fmt.Println(" Chaos Monkey struck! Returning 500 Error.")
			w.WriteHeader(http.StatusInternalServerError)

			w.Write([]byte("Chaos Monkey: Service unavailable!"))

			return
		}

		fmt.Println(" Simulating network lag (2 seconds)...")
		time.Sleep(2 * time.Second)

		next.ServeHTTP(w, r)
	})
}
