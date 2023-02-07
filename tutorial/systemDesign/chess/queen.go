package main

type queen struct {
	white  bool
	killed bool
}

func (this *queen) canMove(board board, currLocation, newLocation location) error {
	return nil
}

func (this *queen) getPieceType() PieceType {
	return Queen
}

func (this *queen) isKilled() bool {
	return this.killed
}

func (this *queen) isWhite() bool {
	return false
}
