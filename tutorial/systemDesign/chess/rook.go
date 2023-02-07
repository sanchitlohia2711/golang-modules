package main

type rook struct {
	white  bool
	killed bool
}

func (this *rook) canMove(board board, currLocation, newLocation location) error {
	return nil
}

func (this *rook) getPieceType() PieceType {
	return Pawn
}

func (this *rook) isKilled() bool {
	return this.killed
}

func (this *rook) isWhite() bool {
	return true
}
