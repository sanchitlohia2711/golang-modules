package main

type GameStatus uint8

const (
	GameInProgress GameStatus = iota
	GameDraw
	Stalemate
	FirstPlayerWin
	SecondPlayerWin
)
