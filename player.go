package main

import "math/rand"

const MODE_DEFENSIVE = 0
const MODE_OFFENSIVE = 1
const MODE_NONE = -1

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

func (p *Player) IsDefensiveHavingNumber(n int) bool {
	for i := 0; i < len(p.DefenseArray); i++ {
		if p.DefenseArray[i] == n {
			return true
		}
	}
	return false

}

func (p *Player) ReverseRole() {
	if p.Mode == MODE_DEFENSIVE {
		p.Mode = MODE_OFFENSIVE
	} else if p.Mode == MODE_OFFENSIVE {
		p.Mode = MODE_DEFENSIVE
	}
}

func (p *Player) SetOffensiveMode() {
	p.Mode = MODE_OFFENSIVE

}

func (p *Player) SetDefensiveMode() {
	p.Mode = MODE_DEFENSIVE
}


func NewPlayer(name string) *Player {
	return &Player{
		Name:         name,
		Mode:         MODE_NONE,
		DefenseArray: nil,
	}
}


func (l *Player) GenerateDefenceArray() {

	if l.Name == "Joey" {
		l.DefenseArray = newDefenceArray(7)

	} else if l.Name == "Monica" {

		l.DefenseArray = newDefenceArray(6)

	} else if l.Name == "Chandler" {

		l.DefenseArray = newDefenceArray(6)

	} else if l.Name == "Ross" {

		l.DefenseArray = newDefenceArray(5)

	} else if l.Name == "Phoebe" {

		l.DefenseArray = newDefenceArray(5)

	} else if l.Name == "Rachel" {

		l.DefenseArray = newDefenceArray(6)

	} else if l.Name == "Sachin" {

		l.DefenseArray = newDefenceArray(4)

	} else if l.Name == "Rohan" {

		l.DefenseArray = newDefenceArray(5)

	} else {

		l.DefenseArray = newDefenceArray(10)
	}



}

func newDefenceArray(arraysize int) []int {

	k := make([]int, arraysize)

	for i := 0; i < arraysize; i++ {
		k[i] = int(rand.Intn(11))
	}

	return k
}



