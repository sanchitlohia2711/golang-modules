package main

func main() {

	board := &board{
		square:    [][]Symbol{{Dot, Dot, Dot}, {Dot, Dot, Dot}, {Dot, Dot, Dot}},
		dimension: 3,
	}

	player1 := &humanPlayer{
		symbol: Cross,
		id:     1,
	}

	player2 := &humanPlayer{
		symbol: Circle,
		id:     2,
	}

	game := initGame(board, player1, player2)
	game.play()
	game.printResult()

	// square := [][]string{{"*", "*", "o"}, {"o", "*", "o"}, {"o", "o", "*"}}
	// b := &board{
	// 	square:    square,
	// 	dimension: 3,
	// }

	// b.printBoard()
	// isWinner, symbol := b.winner()
	// fmt.Println(isWinner)
	// fmt.Println(symbol)

	// square = [][]string{{"*", "o", "*"}, {"o", "*", "o"}, {"o", "*", "o"}}
	// b = &board{
	// 	square:    square,
	// 	dimension: 3,
	// }

	// b.printBoard()
	// isWinner, symbol = b.winner()
	// fmt.Println(isWinner)
	// fmt.Println(symbol)

	// square = [][]string{{"", "", "o"}, {"o", "*", "o"}, {"*", "*", "*"}}
	// b = &board{
	// 	square:    square,
	// 	dimension: 3,
	// }

	// b.printBoard()
	// isWinner, symbol = b.winner()
	// fmt.Println(isWinner)
	// fmt.Println(symbol)

	// square = [][]string{{"", "", "o"}, {"", "o", ""}, {"o", "*", "*"}}
	// b = &board{
	// 	square:    square,
	// 	dimension: 3,
	// }

	// b.printBoard()
	// isWinner, symbol = b.winner()
	// fmt.Println(isWinner)
	// fmt.Println(symbol)
}
