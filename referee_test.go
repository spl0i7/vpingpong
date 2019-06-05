package main

import (
	"testing"
)

func TestNewReferee(t *testing.T) {

	r := NewReferee()
	if r == nil {
		t.Fatal("did not expect referee to be null")
	}
}

func TestReferee_StartGame(t *testing.T) {

	r := NewReferee()

	r.AddPlayer(NewPlayer("A"))
	r.AddPlayer(NewPlayer("B"))
	r.AddPlayer(NewPlayer("C"))
	r.AddPlayer(NewPlayer("D"))
	r.AddPlayer(NewPlayer("E"))
	r.AddPlayer(NewPlayer("F"))
	r.AddPlayer(NewPlayer("G"))
	r.AddPlayer(NewPlayer("H"))

	err := r.StartGame()
	if err != nil {
		t.Fatal(err)
	}

	r = NewReferee()

	r.AddPlayer(NewPlayer("A"))
	r.AddPlayer(NewPlayer("B"))
	r.AddPlayer(NewPlayer("C"))
	r.AddPlayer(NewPlayer("D"))
	err = r.StartGame()
	if err == nil {
		t.Fatal("expected error")
	}







}

func TestReferee_AddPlayer(t *testing.T) {

	r := NewReferee()
	for i := 0; i < 8; i++ {

		err := r.AddPlayer(NewPlayer("A"))
		if err != nil {
			t.Fatal("expected to add at least 8 players")
		}
	}

	r = NewReferee()

	for i := 0; i < 10; i++ {

		err := r.AddPlayer(NewPlayer("A"))

		if  i == 8 && err == nil {
			t.Fatal("did not expect to add more than 8 players")
		}
	}


}

func TestReferee_CanPlayGame(t *testing.T) {

	r := NewReferee()
	for i := 0; i < 7; i++ {

		err := r.AddPlayer(NewPlayer("A"))
		if err != nil {
			t.Fatal("expected to add at least 8 players")
		}

		if r.CanStartGame() {
			t.Fatal("did not expect to start game with less than 8 players")
		}
	}

	_ = r.AddPlayer(NewPlayer("A"))
	if !r.CanStartGame() {
		t.Fatal("expected to start game with 8 players")
	}

}

func TestReferee_RemoveLoser(t *testing.T) {

	p1 := NewPlayer("A")
	p2 := NewPlayer("B")

	r := NewReferee()

	r.AddPlayer(p1)
	r.AddPlayer(p2)

	r.RemoveLoser(p1, p1, p2)

	if  r.Players[0] != p1 {
		t.Fatal("did not remove loser player")
	}

	r = NewReferee()
	r.AddPlayer(p1)
	r.AddPlayer(p2)


	r.RemoveLoser(p2, p1, p2)

	if  r.Players[0] != p2 {
		t.Fatal("did not remove loser player")
	}





}
