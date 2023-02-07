package main

type king struct {
	white  bool
	killed bool
}

func (this *king) canMove(board board, currLocation, newLocation location) error {
	return nil
}

func (this *king) getPieceType() PieceType {
	return King
}

func (this *king) isKilled() bool {
	return this.killed
}

func (this *king) isWhite() bool {
	return this.white
}
