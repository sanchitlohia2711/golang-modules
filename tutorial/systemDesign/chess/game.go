package main

import "fmt"

type game struct {
	board           *board
	firstPlayer     iPlayer
	secondPlayer    iPlayer
	firstPlayerTurn bool
	moveIndex       int
	gameStatus      GameStatus
	moves           []move
}

func initGame(b *board, p1, p2 iPlayer) *game {
	game := &game{
		board:           b,
		firstPlayer:     p1,
		secondPlayer:    p2,
		firstPlayerTurn: true,
		gameStatus:      GameInProgress,
	}
	return game
}

func (this *game) play() error {
	var move move
	var err error
	for {
		if this.firstPlayerTurn {
			move = this.firstPlayer.getNextMove(*this.board)
			if move.resign {
				this.setGameStatus(true, true, false)
			}
			if move.drawOffer {
				if this.secondPlayer.agreeDraw(*this.board) {
					this.setGameStatus(false, true, false)
					break
				}
			}
			err := this.board.makeMove(move)
			if err != nil {
				return err
			}
		} else {
			move = this.firstPlayer.getNextMove(*this.board)
			if move.resign {
				this.setGameStatus(true, true, false)
			}
			if move.drawOffer {
				if this.secondPlayer.agreeDraw(*this.board) {
					this.setGameStatus(false, true, false)
					break
				}
			}
			err = this.board.makeMove(move)
			if err != nil {
				return err
			}
		}
		this.moves = append(this.moves, move)
		win, draw, stalemate := this.checkGameStatus()
		this.setGameStatus(win, draw, stalemate)
		if this.gameStatus != GameInProgress {
			break
		}
	}
	return nil
}

func (this *game) checkGameStatus() (bool, bool, bool) {

	return true, true, true

}

func (g *game) setGameStatus(win bool, whiteWon bool, stalemate bool) {
	if win {
		if whiteWon {
			if g.firstPlayer.isWhite() {
				g.gameStatus = FirstPlayerWin
				return
			} else {
				g.gameStatus = SecondPlayerWin
				return
			}
		}
	}
	if stalemate {
		g.gameStatus = Stalemate
	}
	g.gameStatus = GameDraw
	return
}

func (g *game) printMove(player iPlayer, x, y int) {

}

func (g *game) printResult() {
	switch g.gameStatus {
	case GameInProgress:
		fmt.Println("Game in Progress")
	case GameDraw:
		fmt.Println("Game Drawn")
	case Stalemate:
		fmt.Println("Stalemate")
	case FirstPlayerWin:
		fmt.Println("First Player Win")
	case SecondPlayerWin:
		fmt.Println("Second Player Win")
	default:
		fmt.Println("Invalid Game Status")
	}
	g.board.printBoard()
}
