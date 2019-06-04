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

func TestPlayer_ReverseRole(t *testing.T) {
	p1 := NewPlayer("A", MODE_DEFENSIVE, newDefenceArray(0))
	if p1 == nil {
		t.Fatal("could not create player")
	}

	old := p1.Mode
	p1.ReverseRole()
	if p1.Mode == old {
		t.Fatal("could not reverse role")
	}

	p1.ReverseRole()

	if p1.Mode != old {
		t.Fatal("could not reverse role")
	}
}

func TestPlayer_GenerateOffensiveNumber(t *testing.T) {
	p1 := NewPlayer("A", MODE_OFFENSIVE, newDefenceArray(0))
	if p1 == nil {
		t.Fatal("could not create player")
	}

	for i := 0; i < 1000; i ++ {
		if p1.GenerateOffensiveNumber() > 10 {
			t.Fatal("unexpectedly generated number greater 10")
		}
	}
}
