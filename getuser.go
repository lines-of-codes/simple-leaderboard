package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	body := GetEntryFromBody(w, r)

	if body == nil {
		return
	}

	var data LeaderboardEntry
	row := getUserStmt.QueryRow(body.Name)
	err := row.Scan(&data.Id, &data.Name, &data.Score, &data.Time, &data.Date)

	if err != nil {
		log.Println(err)
		http.Error(w, "An error occurred while executing an SQL statement.", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(data)

	if err != nil {
		http.Error(w, "An error occurred while converting database record to JSON.", http.StatusInternalServerError)
		return
	}
}
