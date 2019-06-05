package main

import "testing"

func TestGame_Play(t *testing.T) {

	for i := 0; i < 1000; i++ {
		p1 := NewPlayer("A")
		p2 := NewPlayer("B")
		p1.SetDefensiveMode()
		p2.SetOffensiveMode()
		g := NewGame(p1, p2)
		winner := g.Play()
		if winner == nil {
			t.Fatal("play should not return null")
		}
	}


}

func TestNewGame(t *testing.T) {

	p1 := NewPlayer("A")
	p2 := NewPlayer("B")

	g := NewGame(p1, p2)

	if g == nil {
		t.Fatal("could not create new game")
	}

	if g.First.Score != 0 || g.Second.Score != 0 {
		t.Fatal("could not initialize game")
	}

}
