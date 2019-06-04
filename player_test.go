package main

import "testing"

func TestNewPlayer(t *testing.T) {
	p1 := NewPlayer("A", MODE_DEFENSIVE, newDefenceArray(0))
	p2 := NewPlayer("B", MODE_OFFENSIVE, newDefenceArray(3))

	if p1 == nil || p2 == nil {
		t.Fatal("could not create player")
	}

	if p1.Name != "A" || p1.Mode != MODE_DEFENSIVE || len(p1.DefenseArray) != 0 {
		t.Fatal("unexpected initialization")
	}

	if p2.Name != "B" || p2.Mode != MODE_OFFENSIVE || len(p2.DefenseArray) != 3 {
		t.Fatal("unexpected initialization")
	}

}
