package main

import (
	"fmt"
)

// Board ...
type Board struct {
	bot      Player
	player   Player
	board    [boardSize]int //don't need to initialize board, as it's all zeroes on creation
	gameOver bool
}

func (b Board) setup(c int) {
	b.bot = b.bot.init(-c, false)
	b.player = b.player.init(c, true)

	b.board[27] = white
	b.board[28] = black
	b.board[35] = black
	b.board[36] = white
}

func (b Board) apply(color, cell int, debug bool) {
	if b.board[cell] == none {
		if debug {
			fmt.Println("Applying move at cell", cell)
			fmt.Println("Cell was originally:", b.board[cell])
		}

		b.board[cell] = color

		if debug {
			fmt.Println("Cell is now:", b.board[cell])
		}
	}
} // end apply()

func (b Board) flipDiscs(color, dir, cell int, debug bool) {
	tempCell := cell

Check:
	for tempCell >= 0 && cell < boardSize {
		tempCell = tempCell + dir

		if debug {
			fmt.Println("Cell is now:", tempCell)
		}

		if tempCell < boardSize {
			if b.board[tempCell] == color {
				break Check
			} else {
				b.board[tempCell] = color
			}
		}
	}
} // end flipDiscs()

func (b Board) print(moveset []Move) {
	// TODO: add in param for printing moves
	// var cells []int
	cells := getCells(moveset)
	// for _, m := range moveset {
	// 	cells = append(cells, m.cell)
	// }

	fmt.Printf("(@) is bot as %v | (O) is player as %v\n", b.bot.color, b.player.color)

	fmt.Println("  ._a_b_c_d_e_f_g_h_")

	for i, e := range b.board {
		if i%8 == 0 {
			fmt.Printf("%v |", getRow(i))
		}

		switch e {
		case none:
			printMove(i, cells)
		case black:
			printBlack(i)
		case white:
			printWhite(i)
		}
	}
}

func (b Board) isGameOver() bool {
	return b.player.passing && b.bot.passing
}
