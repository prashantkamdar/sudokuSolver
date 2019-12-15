package main

import "fmt"

func main() {
	var board [9][9]int
	populateBoard(&board)
	printBoard(board)
}

func populateBoard(board *[9][9]int) {
	for row := 0; row < 9; row++ {
		value := row + 1
		for col := 0; col < 9; col++ {
			board[row][col] = value
		}
	}
}

func printBoard(board [9][9]int) {
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		fmt.Print("| ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}
