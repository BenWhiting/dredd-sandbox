package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func failingTest(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func passingTest(w http.ResponseWriter, req *http.Request) {
	req.Header.Add("content-type", "application/json")
	mapD := map[string]string{"message": ""}
	mapB, _ := json.Marshal(mapD)
	fmt.Fprintf(w, string(mapB))
}

func main() {
	http.HandleFunc("/fail", failingTest)
	http.HandleFunc("/pass", passingTest)
	fmt.Println("starting server on port 7999")
	http.ListenAndServe(":7999", nil)
}
