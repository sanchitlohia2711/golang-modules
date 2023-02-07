package main

type bishop struct {
	white  bool
	killed bool
}

func (this *bishop) canMove(board board, currLocation, newLocation location) error {
	return nil
}

func (this *bishop) getPieceType() PieceType {
	return Pawn
}

func (this *bishop) isKilled() bool {
	return this.killed
}

func (this *bishop) isWhite() bool {
	return this.white
}
