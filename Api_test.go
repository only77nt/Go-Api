package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetValues(t *testing.T) {
	values = append(values, Value{ID: "0", Number: "5"})
	values = append(values, Value{ID: "1", Number: "7"})
	req, err := http.NewRequest("GET", "/api/retrieve/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetValues)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"ID":"0","number":"5"},{"ID":"1","number":"7"}]`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetValue(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/retrieve/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetValue)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"ID":"1","number":"7"}`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateValue(t *testing.T) {

	var jsonStr = []byte(`{"ID":"8081","number":"887"}`)

	req, err := http.NewRequest("POST", "/api/generate/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateValue)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"ID":"8081","number":"887"}`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateValueFormat(t *testing.T) {

	var jsonStr = []byte(`{"ID":"1847","number":"93458"}`)

	req, err := http.NewRequest("POST", "/api/generate/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("type", "int")
	q.Add("len", "5")
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateValue)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"ID":"1847","number":"93458"}`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
