package main

type iPlayer interface {
	isWhite() bool
	getNextMove(board) move
	agreeDraw(board) bool
	getID() int
}
