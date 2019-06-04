package main

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
			Player: p1,
			Score: 0,
		},
	}
}

func (g *Game) Play() {

	// we already know p1 is offensive and p2 is defensive

	for g.First.Score == 5 || g.Second.Score == 5 {

		var offensive, defensive *ScoreBoard
		if g.First.Player.Mode == MODE_OFFENSIVE {
			offensive = g.First
			defensive = g.Second
		} else {
			offensive = g.First
			defensive = g.Second
		}

		offensiveChoice := offensive.Player.GenerateOffensiveNumber()

		if defensive.Player.IsDefensiveHavingNumber(offensiveChoice) {
			defensive.Score += 1
		} else {
			offensive.Score += 1
		}

	}

}
