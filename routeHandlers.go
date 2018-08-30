package main

import (
	"encoding/json"
	"io"
	"net/http"
)

// HandleGetBlockchain returns the full blockchsin as JSON
func HandleGetBlockchain(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(w, string(bytes))
}
