package main

import (
    "encoding/json"
    "net/http/httptest"
    "testing"
)

func TestHealth(t *testing.T) {
    request := httptest.NewRequest("GET", "/health", nil)
    response := httptest.NewRecorder()
    handler().ServeHTTP(response, request)
    var body map[string]string
    if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil { t.Fatal(err) }
    if response.Code != 200 || body["service"] != serviceName { t.Fatalf("unexpected response: %d %v", response.Code, body) }
}