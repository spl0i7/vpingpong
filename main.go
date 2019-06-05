package main

func main() {

	r := NewReferee()

	r.AddPlayer(NewPlayer("Joey"))
	r.AddPlayer(NewPlayer("Monica"))
	r.AddPlayer(NewPlayer("Chandler"))
	r.AddPlayer(NewPlayer("Ross"))
	r.AddPlayer(NewPlayer("Phoebe"))
	r.AddPlayer(NewPlayer("Rachel"))
	r.AddPlayer(NewPlayer("Sachin"))
	r.AddPlayer(NewPlayer("Rohan"))


	r.StartGame()

}
