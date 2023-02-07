package main

type humanPlayer struct {
	white bool
	id    int
}

func (this humanPlayer) getID() int {
	return this.id
}

func (this humanPlayer) isWhite() bool {
	return this.white
}

func (this humanPlayer) getNextMove(board board) move {
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

func (this humanPlayer) agreeDraw(board board) bool {
	return false
}
