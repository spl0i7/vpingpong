package main

import "fmt"

type ScoreBoard struct {
	Player *Player
	Score int
}
type Game struct {
	First *ScoreBoard
	Second *ScoreBoard
}

func NewGame(p1, p2 *Player)  *Game {

	return &Game{
		First: &ScoreBoard{
			Player: p1,
			Score: 0,
		},
		Second: &ScoreBoard{
			Player: p2,
			Score: 0,
		},
	}
}

func (g *Game) Play() *Player {

	for g.First.Score < 5 && g.Second.Score < 5 {

		var offensive, defensive *ScoreBoard

		if g.First.Player.Mode == MODE_OFFENSIVE {
			offensive = g.First
			defensive = g.Second
		} else {
			offensive = g.Second
			defensive = g.First
		}

		offensiveChoice := offensive.Player.GenerateOffensiveNumber()
		fmt.Printf("%s Chose : %d\n", offensive.Player.Name, offensiveChoice)

		defensive.Player.GenerateDefenceArray()
		fmt.Printf("%s has : ", defensive.Player.Name)
		for _, j := range defensive.Player.DefenseArray {
			fmt.Printf("%d ", j)
		}
		fmt.Println()

		if defensive.Player.IsDefensiveHavingNumber(offensiveChoice) {
			defensive.Score += 1
			fmt.Printf("%s Wins\n", defensive.Player.Name)
		} else {
			offensive.Score += 1
			fmt.Printf("%s Wins\n", offensive.Player.Name)
		}

		offensive.Player.ReverseRole()
		defensive.Player.ReverseRole()
	}

	fmt.Printf("\nScore : %s - %d\t%s - %d\n", g.First.Player.Name, g.First.Score, g.Second.Player.Name, g.Second.Score)
	if g.First.Score == 5 {
		return g.First.Player
	}
	return g.Second.Player

}
