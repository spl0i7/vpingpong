package main

import "testing"

func TestGame_Play(t *testing.T) {



}

func TestNewGame(t *testing.T) {

	p1 := NewPlayer("A", newDefenceArray(3))
	p2 := NewPlayer("B", newDefenceArray(4))
	p1.SetDefensiveMode()
	p2.SetDefensiveMode()

	g := NewGame(p1, p2)

	if g == nil {
		t.Fatal("could not create new game")
	}

	if g.First.Score != 0 || g.Second.Score != 0 {
		t.Fatal("could not initialize game")
	}
	

}
