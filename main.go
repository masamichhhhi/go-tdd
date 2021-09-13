package main

import (
	"fmt"
	"log"
	"os"

	"github.com/masamichhhhi/go-tdd/server"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := server.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	game := server.NewTexasHoldem(server.BlindAlerterFunc(server.Alerter), store)
	cli := server.NewCLI(os.Stdin, os.Stdout, game)

	cli.PlayPoker()

	// server := server.NewPlayerServer(store)

	// if err := http.ListenAndServe(":5000", server); err != nil {
	// 	log.Fatalf("could not listen on port 5000 %v", err)
	// }
}
