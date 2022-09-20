package main

import (
	"fmt"
	"math/rand"
	"time"
)


type ticktackBorad [3][3]rune

func main() {
	rand.Seed(time.Now().UnixNano())
	
	var playerMove bool
	var whoWon string
	var win bool

	var board ticktackBorad

	fmt.Println("Starting Game Board empty")
	board.displayBoard()

	if rand.Intn(2) == 0{
		playerMove = true
	}else{
		playerMove = false
	}

	for i:= 0 ; i <9; i++{
		if playerMove{
			fmt.Println("Player Move", i +1)
			board.player()
			playerMove = false
		}else {
			fmt.Println("Computer Move", i +1)
			board.computer()
			playerMove = true
		}

		if whoWon, win = board.check(); win{
			break
		}
		board.displayBoard()
	}


	fmt.Println("*****%v won *** final board view,", whoWon)
}

func (t ticktackBorad) displayBoard(){
	fmt.Print("-------------")
	for i := 0; i < 3; i++{
		fmt.Print("\n|")
		for j := 0; j < 3; j++{
			fmt.Printf("%c |", t[i][j])
		}
		fmt.Print("\n ----------")
	}
	fmt.Print("\n")
}

func (t *ticktackBorad) player(){
	var x, y int
	fmt.Println("Enter the row(1-3) and the column(1-3) positions: ")
	
	if _,err := fmt.Scan(&x, &y); err == nil{
		x--
		y--
		if (x >=0 && x<=2) && (y >= 0 && y<= 2) && (t[x][y] == 0){
			t[x][y] = 'X'
		}else{
			fmt.Println("Invalid input try again.")
			t.player()
		}
	}
}