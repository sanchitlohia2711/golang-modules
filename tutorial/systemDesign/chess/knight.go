package main

type knight struct {
	white  bool
	killed bool
}

func (this *knight) canMove(board board, currLocation, newLocation location) error {
	return nil
}

func (this *knight) getPieceType() PieceType {
	return Knight
}

func (this *knight) isKilled() bool {
	return this.killed
}

func (this *knight) isWhite() bool {
	return this.white
}
