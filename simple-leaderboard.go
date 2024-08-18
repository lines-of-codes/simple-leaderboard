package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// Initialize the database with the tables
// Returns a boolean indicating if there is an error while executing the query.
func initDb(db *sql.DB) {
	log.Println("Initializing leaderboard...")
	sqlStmt := `
  CREATE TABLE leaderboard (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
    name TEXT NOT NULL,
    score INTEGER NOT NULL,
    time INTEGER NOT NULL,
    date INTEGER NOT NULL
  ) STRICT;
  `

	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully initialized leaderboard database!")
}

func prepareStatement(db *sql.DB, sql string) *sql.Stmt {
	stmt, err := db.Prepare(sql)

	if err != nil {
		log.Fatal(err)
	}

	return stmt
}

const topQueryAmount = 10

var insertStmt *sql.Stmt
var getUserStmt *sql.Stmt
var getTopStmt *sql.Stmt

func main() {
	log.Println("Starting the simple leaderboard server!")

	init := flag.Bool("init", false, "If set, Initialize the database before starting the HTTP API server.")
	addr := flag.String("addr", ":5432", "Specifies the address the HTTP server should listen to.")

	flag.Parse()

	db, err := sql.Open("sqlite3", "./leaderboard.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if *init {
		initDb(db)
	}

	insertStmt = prepareStatement(db, "INSERT INTO leaderboard(name, score, time, date) VALUES(?, ?, ?, unixepoch());")
	getUserStmt = prepareStatement(db, "SELECT * FROM leaderboard WHERE name=?;")
	getTopStmt = prepareStatement(db, fmt.Sprintf("SELECT * FROM leaderboard ORDER BY score DESC LIMIT %d", topQueryAmount))

	http.HandleFunc("POST /api/add", AddEntry)
	http.HandleFunc("GET /api/get/user", GetUser)
	http.HandleFunc("GET /api/get/top", GetTop)

	log.Fatal(http.ListenAndServe(*addr, nil))
}
