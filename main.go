package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	serveMux := http.NewServeMux()

	cfg := ApiConfig{fileServerHits: 0}

	serveMux.Handle("/app/*", cfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	serveMux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(http.StatusText(http.StatusOK)))
	})
	serveMux.HandleFunc("GET /metrics", cfg.metricsHandler)
	serveMux.HandleFunc("/reset", cfg.resetHandler)
	server := http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}

	log.Printf("Serving on port: %s\n", port)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
