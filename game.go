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
			Player: p2,
			Score: 0,
		},
	}
}

func (g *Game) Play() *Player {

	for g.First.Score < 5 || g.Second.Score < 5 {

		var offensive, defensive *ScoreBoard

		if g.First.Player.Mode == MODE_OFFENSIVE {
			offensive = g.First
			defensive = g.Second
		} else {
			offensive = g.Second
			defensive = g.First
		}

		offensiveChoice := offensive.Player.GenerateOffensiveNumber()
		defensive.Player.GenerateDefenceArray()
		if defensive.Player.IsDefensiveHavingNumber(offensiveChoice) {
			defensive.Score += 1
		} else {
			offensive.Score += 1
		}
	}

	if g.First.Score == 5 {
		return g.First.Player
	}
	return g.Second.Player

}
