package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

type (
	Report struct {
		Type string `json:"reportType"`
	}
)

func createReportJSONParser(w http.ResponseWriter, r *http.Request) {
	var report Report

	var jSON = jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
		CaseSensitive:          true,
	}.Froze()

	err := jSON.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		panic(err)
	}
	fmt.Println(report)
}

func createReport(w http.ResponseWriter, r *http.Request) {
	var report Report

	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		panic(err)
	}

	fmt.Println(report)

	json.NewEncoder(w).Encode(report)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/person/create", createReport)
	mux.HandleFunc("/report/create", createReportJSONParser)

	log.Println("Starting server on :4000...")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
