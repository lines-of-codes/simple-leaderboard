package main

import (
	"net/http"
)

func AddEntry(w http.ResponseWriter, r *http.Request) {
	body := GetEntryFromBody(w, r)

	if body == nil {
		return
	}

	_, err := insertStmt.Exec(body.Name, body.Score, body.Time)

	if err != nil {
		http.Error(w, "An error occurred during executing an SQL statement.", http.StatusInternalServerError)
		return
	}
}
