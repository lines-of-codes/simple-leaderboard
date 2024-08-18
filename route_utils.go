package main

import (
	"encoding/json"
	"io"
	"net/http"
)

func GetEntryFromBody(w http.ResponseWriter, r *http.Request) *LeaderboardEntry {
	rawBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "An error occurred while reading the request body.", http.StatusInternalServerError)
		return nil
	}
	defer r.Body.Close()

	var body LeaderboardEntry
	err = json.Unmarshal(rawBody, &body)

	if err != nil {
		http.Error(w, "An error occurred while parsing the JSON in request body.", http.StatusBadRequest)
		return nil
	}

	return &body
}
