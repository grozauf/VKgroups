package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func routes() http.Handler {
	request := http.NewServeMux()
	request.Handle("/auth", http.HandlerFunc(auth))
	return request
}

func auth(writer http.ResponseWriter, _ *http.Request) {

	resp, err := http.Get("https://oauth.vk.com/authorize?client_id=7897817&display=page&redirect_uri=https://oauth.vk.com/blank.html&scope=groups&response_type=token&v=5.52")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Got response: %v", resp)

	// headers
	for name, values := range resp.Header {
		writer.Header()[name] = values
	}

	// status (must come after setting headers and before copying body)
	writer.WriteHeader(resp.StatusCode)

	// body
	io.Copy(writer, resp.Body)
}

func main() {
	handlers := http.NewServeMux()
	handlers.Handle("/", routes())
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      handlers,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Printf("ListenAndServe failed: %v", err)
	}
}
