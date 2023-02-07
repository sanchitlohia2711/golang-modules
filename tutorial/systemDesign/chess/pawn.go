package main

type pawn struct {
	white  bool
	killed bool
}

func (this *pawn) canMove(board board, currLocation, newLocation location) error {
	return nil
}

func (this *pawn) getPieceType() PieceType {
	return Pawn
}

func (this *pawn) isKilled() bool {
	return this.killed
}

func (this *pawn) isWhite() bool {
	return true
}
