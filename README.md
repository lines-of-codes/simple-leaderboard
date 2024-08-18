# simple-leaderboard
A simple self-hostable leaderboard &amp; score database.

This is written in Go with [SQLite](https://github.com/mattn/go-sqlite3)

## Setup

1. Build/Download server executable
2. Initialize database
3. Uh... Done!

The leaderboard database doesn't exist yet, and you'll have to create it.
You can do that by executing `simple-leaderboard -init`. And after that,
You can just launch the server again by simply calling `simple-leaderboard`

By default, the server listens on port `:5432`, If you don't like that, 
You can provide the `-addr` flag with the address you want the server to 
listen on, such as `-addr ":8000"` to listen on port 8000.

### Build it yourself

#### Prerequisites
- git
- go
- gcc (required for sqlite3)

To build the leaderboard server, 
Clone the git repository: `git clone https://github.com/lines-of-codes/simple-leaderboard.git`

Then in the project folder, you can simply run `go build` to build the project.

## License
This software is licensed under the GNU GPLv3.

