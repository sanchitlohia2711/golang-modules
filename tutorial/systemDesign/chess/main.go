package main

func main() {

	whiteKing, whiteQueen, whiteRooks, whiteBishops, whiteKnights, whitePawns := makePieces(true)
	blackKing, blackQueen, blackRooks, blackBishops, blackKnights, blackPawns := makePieces(true)

	cells := make([][]cell, 8)

	for i := 0; i < 8; i++ {
		cells[i] = make([]cell, 8)
	}

	//Fill White Pieces in the first row
	cells[0][0] = cell{location: location{i: 0, j: 0}, piece: whiteRooks[0]}
	cells[0][1] = cell{location: location{i: 0, j: 1}, piece: whiteKnights[0]}
	cells[0][2] = cell{location: location{i: 0, j: 2}, piece: whiteBishops[0]}
	cells[0][3] = cell{location: location{i: 0, j: 3}, piece: whiteKing}
	cells[0][4] = cell{location: location{i: 0, j: 4}, piece: whiteQueen}
	cells[0][5] = cell{location: location{i: 0, j: 5}, piece: whiteBishops[1]}
	cells[0][6] = cell{location: location{i: 0, j: 6}, piece: whiteKnights[1]}
	cells[0][7] = cell{location: location{i: 0, j: 7}, piece: whiteRooks[1]}
	//Fill White Pawns in the second row
	for i := 0; i < 8; i++ {
		cells[1][i] = cell{location: location{i: 0, j: 7}, piece: whitePawns[i]}
	}

	//Fill Black Pieces in the first row
	cells[7][0] = cell{location: location{i: 7, j: 0}, piece: blackRooks[0]}
	cells[7][1] = cell{location: location{i: 7, j: 1}, piece: blackKnights[0]}
	cells[7][2] = cell{location: location{i: 7, j: 2}, piece: blackBishops[0]}
	cells[7][3] = cell{location: location{i: 7, j: 3}, piece: blackKing}
	cells[7][4] = cell{location: location{i: 7, j: 4}, piece: blackQueen}
	cells[7][5] = cell{location: location{i: 7, j: 5}, piece: blackBishops[1]}
	cells[7][6] = cell{location: location{i: 7, j: 6}, piece: blackKnights[1]}
	cells[7][7] = cell{location: location{i: 7, j: 7}, piece: blackRooks[1]}
	//Fill Black Pawns in the second row
	for i := 0; i < 8; i++ {
		cells[6][i] = cell{location: location{i: 0, j: 7}, piece: blackPawns[i]}
	}

	board := &board{
		square:    cells,
		dimension: 8,
	}

	player1 := humanPlayer{
		white: true,
		id:    1,
	}

	player2 := computerPlayer{
		white: false,
		id:    1,
	}

	game := initGame(board, player1, player2)
	game.play()
	game.printResult()
}

func makePieces(isWhite bool) (*king, *queen, [2]*rook, [2]*bishop, [2]*knight, [8]*pawn) {

	king := &king{
		white: isWhite,
	}

	queen := &queen{
		white: isWhite,
	}

	rooks := [2]*rook{}
	for i := 0; i < 2; i++ {
		r := &rook{
			white: true,
		}
		rooks[i] = r
	}

	bishops := [2]*bishop{}
	for i := 0; i < 2; i++ {
		b := &bishop{
			white: true,
		}
		bishops[i] = b
	}

	knights := [2]*knight{}
	for i := 0; i < 2; i++ {
		k := &knight{
			white: true,
		}
		knights[i] = k
	}

	pawns := [8]*pawn{}
	for i := 0; i < 8; i++ {
		p := &pawn{
			white: isWhite,
		}
		pawns[i] = p
	}

	return king, queen, rooks, bishops, knights, pawns
}
