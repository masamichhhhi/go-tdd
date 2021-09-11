package server_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/masamichhhhi/go-tdd/server"
)

var dummyBlindAlerter = &SpyBlindAlerter{}
var dummyPlayerStore = &server.StubPlayerStore{}
var dummyStdIn = &bytes.Buffer{}
var dummyStdOut = &bytes.Buffer{}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type SpyBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *SpyBlindAlerter) ScheduledAlertAt(at time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{at, amount})
}

type GameSpy struct {
	StartCalledWith  int
	FinishCalledWith string
}

func (g *GameSpy) Start(numberOfPlayers int) {
	g.StartCalledWith = numberOfPlayers
}

func (g *GameSpy) Finish(winner string) {
	g.FinishCalledWith = winner
}

func TestCLI(t *testing.T) {
	// t.Run("record chris win from user input", func(t *testing.T) {
	// 	in := strings.NewReader("Chris wins\n")
	// 	playerStore := &server.StubPlayerStore{}

	// 	cli := server.NewCLI(playerStore, in, dummySpyAlerter)
	// 	cli.PlayerPoker()

	// 	server.AssertPlayerWin(t, playerStore, "Chris")
	// })

	// t.Run("record cleo win from user input", func(t *testing.T) {
	// 	in := strings.NewReader("Cleo wins\n")
	// 	playerStore := &server.StubPlayerStore{}

	// 	cli := server.NewCLI(playerStore, in, dummySpyAlerter)
	// 	cli.PlayerPoker()

	// 	server.AssertPlayerWin(t, playerStore, "Cleo")
	// })

	// t.Run("it scheduls printing of blind values", func(t *testing.T) {
	// 	in := strings.NewReader("Chris wins\n")
	// 	playerStore := &server.StubPlayerStore{}
	// 	blindAlerter := &SpyBlindAlerter{}

	// 	cli := server.NewCLI(playerStore, in, blindAlerter)
	// 	cli.PlayerPoker()

	// 	cases := []struct {
	// 		expectedScheduleTime time.Duration
	// 		expectedAmount       int
	// 	}{
	// 		{0 * time.Second, 100},
	// 		{10 * time.Minute, 200},
	// 		{20 * time.Minute, 300},
	// 		{30 * time.Minute, 400},
	// 		{40 * time.Minute, 500},
	// 		{50 * time.Minute, 600},
	// 		{60 * time.Minute, 800},
	// 		{70 * time.Minute, 1000},
	// 		{80 * time.Minute, 2000},
	// 		{90 * time.Minute, 4000},
	// 		{100 * time.Minute, 8000},
	// 	}

	// 	for i, c := range cases {
	// 		t.Run(fmt.Sprintf("%d scheduled for %v", c.expectedAmount, c.expectedScheduleTime), func(t *testing.T) {
	// 			if len(blindAlerter.alerts) <= i {
	// 				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.alerts)
	// 			}

	// 			alert := blindAlerter.alerts[i]

	// 			amountGot := alert.amount
	// 			if amountGot != c.expectedAmount {
	// 				t.Errorf("got amount %d, want %d", amountGot, c.expectedAmount)
	// 			}

	// 			gotScheduledTime := alert.at
	// 			if gotScheduledTime != c.expectedScheduleTime {
	// 				t.Errorf("got scheduled time of %v, want %v", gotScheduledTime, c.expectedScheduleTime)
	// 			}
	// 		})
	// 	}
	// })

	t.Run("it prompts the user to enter the number of players", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("7\n")
		game := &GameSpy{}

		cli := server.NewCLI(in, stdout, game)
		cli.PlayPoker()

		gotPrompt := stdout.String()
		wantPrompt := server.PlayerPrompt

		if gotPrompt != wantPrompt {
			t.Errorf("got %q, want %q", gotPrompt, wantPrompt)
		}

		if game.StartCalledWith != 7 {
			t.Errorf("wanted Start called with 7 but got %d", game.StartCalledWith)
		}
	})
}

func assertScheduledAlert(t *testing.T, got, want server.ScheduledAlert) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
