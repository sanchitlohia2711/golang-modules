package main

type PieceType uint8

const (
	King PieceType = iota
	Queen
	Rook
	Bishop
	Knight
	Pawn
)
