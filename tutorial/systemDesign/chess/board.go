package main

type board struct {
	square    [][]cell
	dimension int
}

func (this *board) printBoard() {

}

func (this *board) makeMove(move move) error {
	err := move.piece.canMove(*this, move.currLocation, move.newLocation)
	if err != nil {
		return err
	}

	//Mark the piece at new location as as the current piece
	this.square[move.newLocation.i][move.newLocation.j].piece = move.piece

	//Mark the piece at current location as nil
	this.square[move.currLocation.i][move.currLocation.j].piece = nil
	return nil
}
