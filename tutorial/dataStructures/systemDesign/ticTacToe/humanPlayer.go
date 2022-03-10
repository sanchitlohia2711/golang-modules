package main

import "fmt"

var (
	MovesPlayer1 = [4][2]int{{1, 1}, {2, 0}, {2, 2}, {2, 1}}
	MovesPlayer2 = [4][2]int{{1, 2}, {0, 2}, {0, 0}, {0, 0}}
	//o o
	//*o
	//***

	//MovesPlayer1 = [5][2]int{{1, 1}, {0, 0}, {2, 1}, {1, 0}, {0, 2}}
	//MovesPlayer2 = [4][2]int{{2, 2}, {2, 0}, {0, 1}, {1, 2}}
	//*o*
	//**o
	//o*o

	//MovesPlayer1 = [3][2]int{{1, 1}, {0, 0}, {0, 1}}
	//MovesPlayer2 = [3][2]int{{0, 2}, {2, 2}, {1, 2}}
	//**o
	//*o
	//o

	//MovesPlayer1 = [3][2]int{{0, 0}, {0, 1}, {0, 2}}
	//MovesPlayer2 = [2][2]int{{1, 0}, {1, 1}}
	//***
	//oo
	//

	//MovesPlayer1 = [3][2]int{{1, 0}, {1, 1}, {1, 2}}
	//MovesPlayer2 = [2][2]int{{2, 0}, {2, 1}}
	//
	//***
	//oo

	//MovesPlayer1 = [3][2]int{{1, 0}, {1, 1}, {0, 2}}
	//MovesPlayer2 = [3][2]int{{2, 0}, {2, 1}, {2, 2}}
	//  *
	//**
	//ooo
)

type humanPlayer struct {
	symbol Symbol
	index  int
	id     int
}

func (h *humanPlayer) getSymbol() Symbol {
	return h.symbol
}

func (h *humanPlayer) getNextMove() (int, int, error) {
	if h.symbol == Cross {
		h.index = h.index + 1
		return MovesPlayer1[h.index-1][0], MovesPlayer1[h.index-1][1], nil
	} else if h.symbol == Circle {
		h.index = h.index + 1
		return MovesPlayer2[h.index-1][0], MovesPlayer2[h.index-1][1], nil
	}
	return 0, 0, fmt.Errorf("Invalid Symbol")
}

func (h *humanPlayer) getID() int {
	return h.id
}
