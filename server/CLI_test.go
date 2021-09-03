package server_test

import (
	"strings"
	"testing"

	"github.com/masamichhhhi/go-tdd/server"
)

func TestCLI(t *testing.T) {
	t.Run("record chris win from user input", func(t *testing.T) {
		in := strings.NewReader("Chris wins\n")
		playerStore := &server.StubPlayerStore{}

		cli := server.NewCLI(playerStore, in)
		cli.PlayerPoker()

		server.AssertPlayerWin(t, playerStore, "Chris")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins\n")
		playerStore := &server.StubPlayerStore{}

		cli := server.NewCLI(playerStore, in)
		cli.PlayerPoker()

		server.AssertPlayerWin(t, playerStore, "Cleo")
	})

}
