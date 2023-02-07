package main

type computerPlayer struct {
	white bool
	id    int
}

func (this computerPlayer) isWhite() bool {
	return this.white
}
func (this computerPlayer) getID() int {
	return this.id
}

func (this computerPlayer) getNextMove(board board) move {
	currLocation := location{
		i: 1,
		j: 0,
	}

	newLocation := location{
		i: 2,
		j: 0,
	}
	return move{
		currLocation: currLocation,
		piece:        board.square[1][0].piece,
		newLocation:  newLocation,
	}
}

func (this computerPlayer) agreeDraw(board board) bool {
	return false
}
