package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetTop(w http.ResponseWriter, r *http.Request) {
	rows, err := getTopStmt.Query()

	if err != nil {
		http.Error(w, "An error occurred while executing an SQL statement.", http.StatusInternalServerError)
		return
	}

	result := make([]LeaderboardEntry, topQueryAmount)

	for i := 0; rows.Next(); i++ {
		rows.Scan(&result[i].Id, &result[i].Name, &result[i].Score, &result[i].Time, &result[i].Date)
	}

	json, err := json.Marshal(result)

	if err != nil {
		http.Error(w, "An error occurred while converting database records to JSON.", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s", json)
}
