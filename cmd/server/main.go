package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
)

const serviceName = "visitor_dispute"

func handler() http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
        w.Header().Set("Content-Type", "application/json; charset=utf-8")
        _ = json.NewEncoder(w).Encode(map[string]string{"status": "ok", "service": serviceName})
    })
    return mux
}

func main() {
    host, port := os.Getenv("HOST"), os.Getenv("PORT")
    if host == "" { host = "0.0.0.0" }
    if port == "" { port = "8080" }
    if err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), handler()); err != nil { panic(err) }
}