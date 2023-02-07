package main

type piece interface {
	canMove(board, location, location) error
	getPieceType() PieceType
	isKilled() bool
	isWhite() bool
}
