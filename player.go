package main

import "math/rand"

const MODE_DEFENSIVE = 0
const MODE_OFFENSIVE = 1

type Player struct {
	Name string
	Mode int
	DefenseArray []int
}


func (p *Player) GenerateOffensiveNumber() int {
	if p.Mode == MODE_DEFENSIVE {
		panic("cannot generate numbers while in offensive mode")
	}
	return rand.Intn(11)
}

func (p *Player) ReverseRole() {
	if p.Mode == MODE_DEFENSIVE {
		p.Mode = MODE_OFFENSIVE
	} else if p.Mode == MODE_OFFENSIVE {
		p.Mode = MODE_DEFENSIVE
	}
}


func NewPlayer(name string, mode int, arr []int) *Player {
	return &Player{
		Name:         name,
		Mode:         mode,
		DefenseArray: arr,
	}
}


func newDefenceArray(arraysize int) []int {

	k := make([]int, arraysize)

	for i := 0; i < arraysize; i++ {
		k[i] = int(rand.Intn(11))
	}

	return k
}



