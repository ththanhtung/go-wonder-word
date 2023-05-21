package main

import (
	"time"

	"github.com/gorilla/websocket"
)

type Player struct {
	userID   string
	username string
	client   *Client
	score    int
	guess    string
}

func NewPlayer(id, name string, conn *websocket.Conn) *Player {
	return &Player{
		userID:   id,
		username: name,
		client:   NewClient(conn),
		score:    0,
		guess:    "",
	}
}

func (p *Player) Update(players []*Player, state *GameState) {
	scoreForOneGuess, _ := state.wonderWordGame.ScoreCalculator(p.guess)
	p.score += scoreForOneGuess
}

func (p *Player) MakeTurn() bool {
	timer := time.NewTimer(10 * time.Second)

	select {
	case <-timer.C:
		return false
	case <-p.client.msgIn:
		return true
	}
}
