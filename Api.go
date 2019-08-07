package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Value struct {
	ID     string `json:"ID"`
	Number string `json:"number"`
}

var values []Value

func GetValues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(values)
}

func GetValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()
	for _, item := range values {
		if item.ID == query["id"][0] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Value{})
}

func CreateValue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var value Value
	_ = json.NewDecoder(r.Body).Decode(&value)
	query := r.URL.Query()
	id := strconv.Itoa(rand.Intn(10000))
	for _, v := range values {
		if v.ID == id {
			return
		}
	}
	if len(query) != 0 {
		t := query["type"][0]
		l, _ := strconv.Atoi(query["len"][0])
		switch {
		case t == "string":
			value.Number = RandString(l)
		case t == "int":
			d := math.Pow10(l) - 1
			value.Number = strconv.Itoa(rand.Intn(int(d)))
		}
	} else {
		value.Number = strconv.Itoa(rand.Intn(int(1000)))
	}
	value.ID = id
	values = append(values, value)
	json.NewEncoder(w).Encode(value)
	return
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func RandString(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func main() {
	fmt.Println("Сервер запущен!")
	r := mux.NewRouter()
	values = append(values, Value{ID: "0", Number: "5"})
	values = append(values, Value{ID: "1", Number: "7"})
	r.HandleFunc("/api/retrieve", GetValues).Methods("GET")
	r.HandleFunc("/api/retrieve/", GetValue).Methods("GET")
	r.HandleFunc("/api/generate/", CreateValue).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
