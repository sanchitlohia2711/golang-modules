package main

import "fmt"

type board struct {
	square    [][]Symbol
	dimension int
}

func (b *board) printBoard() {
	for i := 0; i < b.dimension; i++ {
		for j := 0; j < b.dimension; j++ {
			if b.square[i][j] == Dot {
				fmt.Print(".")
			} else if b.square[i][j] == Cross {
				fmt.Print("*")
			} else {
				fmt.Print("o")
			}

		}
		fmt.Println("")
	}
}

func (b *board) markSymbol(i, j int, symbol Symbol) (bool, Symbol, error) {
	if i > b.dimension || j > b.dimension {
		return false, Dot, fmt.Errorf("index input is greater than dimension")
	}
	if b.square[i][j] != Dot {
		return false, Dot, fmt.Errorf("input square already marked")
	}
	if symbol != Cross && symbol != Circle {
		return false, Dot, fmt.Errorf("incorrect Symbol")
	}
	b.square[i][j] = symbol
	win := b.checkWin(i, j, symbol)

	return win, symbol, nil
}

func (b *board) checkWin(i, j int, symbol Symbol) bool {
	//Check Row
	rowMatch := true
	for k := 0; k < b.dimension; k++ {
		if b.square[i][k] != symbol {
			rowMatch = false
		}
	}

	if rowMatch {
		return rowMatch
	}

	//Check Row
	columnMatch := true
	for k := 0; k < b.dimension; k++ {
		if b.square[k][j] != symbol {
			columnMatch = false
		}
	}

	if columnMatch {
		return columnMatch
	}

	//Check diagonal
	diagonalMatch := false
	if i == j {
		diagonalMatch = true
		for k := 0; k < b.dimension; k++ {
			if b.square[k][k] != symbol {
				diagonalMatch = false
			}
		}
	}

	return diagonalMatch

}

func (b *board) checkWin2() (bool, Symbol) {
	win, symbol := b.rowCrossed()
	if win {
		return true, symbol
	}

	win, symbol = b.columnCrossed()
	if win {
		return true, symbol
	}

	win, symbol = b.diagonalCrossed()
	if win {
		return true, symbol
	}

	return false, Dot
}

func (b *board) rowCrossed() (bool, Symbol) {
	rowCrossed := false
	var firstSymbol Symbol
	for i := 0; i < b.dimension; i++ {
		firstSymbol = b.square[i][0]
		if b.square[i][0] == Dot {
			continue
		}
		j := 1
		rowCrossed = true
		for j = 1; j < b.dimension; j++ {
			if b.square[i][j] != firstSymbol {
				rowCrossed = false
				break
			}
		}
		if rowCrossed {
			break
		}

	}
	return rowCrossed, firstSymbol
}

func (b *board) columnCrossed() (bool, Symbol) {
	columnCrossed := false
	var firstSymbol Symbol
	for i := 0; i < b.dimension; i++ {
		firstSymbol = b.square[0][i]
		if b.square[0][i] == Dot {
			continue
		}
		j := 1
		columnCrossed = true
		for j = 1; j < b.dimension; j++ {
			if b.square[j][i] != firstSymbol {
				columnCrossed = false
				break
			}
		}
		if columnCrossed {
			break
		}

	}
	return columnCrossed, firstSymbol
}

func (b *board) diagonalCrossed() (bool, Symbol) {
	win, symbol := b.checkFirstDiagonal()
	if win {
		return true, symbol
	}

	win, symbol = b.checkSecondDiagonal()
	if win {
		return true, symbol
	}
	return false, Dot
}

func (b *board) checkFirstDiagonal() (bool, Symbol) {
	if b.square[0][0] == Dot {
		return false, Dot
	}

	firstSymbol := b.square[0][0]

	i := 1
	diagonalCrossed := true
	for i = 1; i < b.dimension; i++ {
		if b.square[i][i] != firstSymbol {
			diagonalCrossed = false
			break
		}
	}

	return diagonalCrossed, firstSymbol
}

func (b *board) checkSecondDiagonal() (bool, Symbol) {
	if b.square[0][b.dimension-1] == Dot {
		return false, Dot
	}

	firstSymbol := b.square[0][b.dimension-1]

	i := 1
	diagonalCrossed := true
	for i = 1; i < b.dimension; i++ {
		if b.square[i][b.dimension-i-1] != firstSymbol {
			diagonalCrossed = false
			break
		}
	}
	return diagonalCrossed, firstSymbol
}
