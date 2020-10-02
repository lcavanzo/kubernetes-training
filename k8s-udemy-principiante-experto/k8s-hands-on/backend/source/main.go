package main

import (
    "net/http"
    "fmt"
    "os"
    "time"
)

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    resp := fmt.Sprintf("La hora es %v y hostname es %v", time.Now(), os.Getenv("Hostname"))
    w.Write([]byte(resp))
}

func main() {
    http.HandleFunc("/", ServeHTTP)
    http.ListenAndServe(":9090", nil)
}
