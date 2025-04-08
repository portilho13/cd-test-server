package main

import (
	"io"
	"log"
	"net/http"
)

func proxyHandler(w http.ResponseWriter, r *http.Request) {

    targetURL := "http://127.0.0.1:1338"
    client := &http.Client{
        CheckRedirect: func(req *http.Request, via []*http.Request) error {
            return http.ErrUseLastResponse
        },
    }

	req, err := http.NewRequest(r.Method, targetURL, r.Body)
    if err != nil {
        http.Error(w, "Error creating request", http.StatusInternalServerError)
        return
    }

    // Copy ALL headers from the original request
    for name, values := range r.Header {
        // Skip some headers that shouldn't be forwarded
        if name == "Connection" || name == "Upgrade" {
            continue
        }
        for _, value := range values {
            req.Header.Add(name, value)
        }
    }


    resp, err := client.Do(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // Remove redirect headers
    resp.Header.Del("Location")
    resp.Header.Del("Refresh")

    for key, values := range resp.Header {
        for _, value := range values {
            w.Header().Add(key, value)
        }
    }

    // Force status 200 OK
    w.WriteHeader(http.StatusOK)

    // Copy body
    _, err = io.Copy(w, resp.Body)
    if err != nil {
        log.Printf("Failed to copy body: %v", err)
    }
}