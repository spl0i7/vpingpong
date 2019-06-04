package main

import (
	"log"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	p1 := NewPlayer("A", newDefenceArray(0))
	p2 := NewPlayer("B", newDefenceArray(3))

	if p1 == nil || p2 == nil {
		t.Fatal("could not create player")
	}

	if p1.Name != "A" || p1.Mode != MODE_NONE|| len(p1.DefenseArray) != 0 {
		t.Fatal("unexpected initialization")
	}

	if p2.Name != "B" || p2.Mode != MODE_NONE || len(p2.DefenseArray) != 3 {
		t.Fatal("unexpected initialization")
	}
}

func TestPlayer_ReverseRole(t *testing.T) {
	p1 := NewPlayer("A", newDefenceArray(0))
	p1.SetDefensiveMode()
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
	p1 := NewPlayer("A", newDefenceArray(0))
	if p1 == nil {
		t.Fatal("could not create player")
	}

	for i := 0; i < 1000; i ++ {
		if p1.GenerateOffensiveNumber() > 10 {
			t.Fatal("unexpectedly generated number greater 10")
		}
	}
}

func TestPlayer_IsDefensiveHavingNumber(t *testing.T) {
	p1 := NewPlayer("A", []int {3, 10, 12 , 134})
	p1.SetDefensiveMode()

	if !p1.IsDefensiveHavingNumber(12) {
		log.Fatal("expected 12 in array")
	}
	if !p1.IsDefensiveHavingNumber(10) {
		log.Fatal("expected 10 in array")
	}
	if p1.IsDefensiveHavingNumber(112) {
		log.Fatal("did not expected 112 in array")
	}
	if p1.IsDefensiveHavingNumber(31) {
		log.Fatal("did not expected 31 in array")
	}
}
