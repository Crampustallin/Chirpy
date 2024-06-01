package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	serveMux := http.NewServeMux()
	serveMux.Handle("/app/*", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	serveMux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(http.StatusText(http.StatusOK)))
	})
	server := http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}

	log.Printf("Serving on port: %s\n", port)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}