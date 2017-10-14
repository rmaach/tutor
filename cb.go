package main

import "fmt"

// rule will contain all the legal moves
var wins int
var rules [][]int


func makerules() {
initrules :=  [][]int{{1,3,6},{1,2,4},{2,4,7},{2,5,9},{3,5,8},{3,6,10},{4,2,1},{4,5,6},
{4,7,11},{4,8,13},{5,8,12},{5,9,14},{6,9,13},{6,10,15},{6,5,4},{6,3,1},{7,4,2},
{7,8,9},{8,5,3},{8,9,10},{9,5,2},{9,8,7},{10,9,8},{10,6,3},{11,7,4},{11,12,13},
{12,8,5},{12,13,14},{13,12,11},{13,14,15},{13,8,4},{13,9,6},{14,13,12},{14,9,5},
{15,10,6},{15,14,13}}

rules = make([][]int, 36)
copy(rules,initrules)
}
// 


func crack(game []int,hist string)  {
	// pegs is used count pegs
	var pegs, i  int

	// used to create a fresh copy of gameboard for each possible move
	var localgame []int 

	// Used to make a copy of history for each possible move
	var localhist string

	// count the pegs that are left
	for i = 1; i <= 15; i++ {
		if (game[i] == 1) {
			pegs++ }
	}

	// test to see if we have won
	if (pegs < 2) {
		wins++
		fmt.Println(hist,"Win# ",wins)
	} else {
		// loop through the legal moves
		for i = 0; i <= 35; i++ {
			if (game[rules[i][0]] == 1 && game[rules[i][1]] == 1 && game[rules[i][2]] == 0) {
				// copy the game
				localgame = make([]int, 16)
				copy(localgame,game)
				// implement the new move
				localgame[rules[i][0]] = 0
				localgame[rules[i][1]] = 0
				localgame[rules[i][2]] = 1
				// format the history 
				localhist = fmt.Sprintf("%v:%v %v %v",hist,rules[i][0], rules[i][1],rules[i][2])
				// Call recursive
				crack(localgame,localhist)
		}
	}
}
}


func main() {
	//--rule := make([]int,16)
	game := []int{}
	// The game board.  We are not using element zero.
	game = []int{99,0,1,1, 1,1,1,1, 1,1,1,1, 1,1,1,1}
	var hist string = "Answer" 

	// Record all the legal moves
	makerules()
	
	fmt.Println(rules)
	fmt.Println(game)
	fmt.Println("Start")
	crack(game,hist)

}
