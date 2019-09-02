package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bionikspoon/learn-go-with-tests/poker"
)

const dbFileName = "game.db.json"

func main() {
	database, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("could not open file err: %#+v\n", err)
	}

	server := poker.NewPlayerServer(poker.NewFileSystemPlayerStore(database))

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("Could not listen on port 5000 %v", err)
	}
}
