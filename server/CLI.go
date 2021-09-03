package server

import "io"

type CLI struct {
	playerStore PlayerStore
	in          io.Reader
}

func (cli *CLI) PlayerPoker() {
	cli.playerStore.RecordWin("Chris")
}
