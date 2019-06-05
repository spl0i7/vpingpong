package main

import (
	"errors"
	"fmt"
)

type Referee struct {
	Players []*Player
}

func NewReferee() *Referee {
	return &Referee{
		Players:nil,
	}
}

func (r *Referee) CanStartGame() bool {
	return len(r.Players) == 8
}

func (r *Referee) AddPlayer(p *Player) error {

	if len(r.Players) < 8 {
		r.Players = append(r.Players, p)
		return nil
	}

	return errors.New("PlayerFull")

}

func (r *Referee) StartGame() error {
	if !r.CanStartGame() {
		return errors.New("NotEnoughPlayers")
	}

	round := 0

	for len(r.Players) > 1 {

		round += 1
		fmt.Printf("=== Round %d ===\n\n", round)

		for i := 0; i < len(r.Players); i += 2 {

			p1 := r.Players[i]
			p1.SetOffensiveMode()

			p2 := r.Players[i+1]
			p2.SetDefensiveMode()

			fmt.Printf("== Match %s vs %s ==\n", p1.Name, p2.Name )
			fmt.Printf("Offensive : %s\nDefensive : %s\n\n", p1.Name, p2.Name )


			game := NewGame(p1, p2)

			winner := game.Play()

			fmt.Printf("== Winner : %s ==\n\n", winner.Name)

			r.RemoveLoser(winner, p1, p2)

		}
	}

	return nil
}

func (r *Referee) RemoveLoser(w, p1, p2 *Player) {

	var loser *Player
	loserIndex := -1

	if w == p1 {
		loser = p2
	} else {
		loser = p1
	}


	for i, p := range r.Players {
		if p == loser {
			loserIndex = i
			break
		}
	}

	lastIndex := len(r.Players) - 1

	r.Players[loserIndex] = r.Players[lastIndex]
	r.Players = r.Players[:lastIndex]

}