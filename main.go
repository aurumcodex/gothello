package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// need to get and parse cli arguments
	const human = true
	const debug = true

	turnCount := 0
	game := Board{}

	reader := bufio.NewReader(os.Stdin)
GetInput:
	for true {
		fmt.Println("what color do you want to play as? (black or white)")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "unable to read user input")
			os.Exit(2)
		}
		input = strings.Replace(input, "\n", "", -1)

		switch input {
		case "b", "B", "Black", "black", "BLACK":
			fmt.Println("player will be set up as black")
			game.setup(black)
			break GetInput
		case "w", "W", "White", "white", "WHITE":
			fmt.Println("player will be set up as white")
			game.setup(white)
			break GetInput
		default:
			fmt.Println("unable to get acceptable input; please re-enter")
			continue GetInput
		}
	}

	currentPlayer := game.player.color
	fmt.Println(game)
	fmt.Println("color: ", color(currentPlayer, true))
	os.Exit(0)

	for !game.gameOver {
		movelist := make([]Move, 0)
		cells := make([]int, 0)

		fmt.Printf("turn count :: %v\n", turnCount)

		if currentPlayer == black {
			movelist = nil
			cells = nil
			movelist = game.generateMoves(game.player.color)

			fmt.Println("legal moves:")
			for _, m := range movelist {
				fmt.Printf("%v %v %v | ", color(game.player.color, false), getCol(m.cell), getRow(m.cell))
				fmt.Printf("num filps: %v | ", m.numFlips)
				fmt.Printf("direction: %v\n", getDir(m.direction))
				cells = append(cells, m.cell)
			}

			if len(movelist) == 0 {
				fmt.Println("player has to pass")
				game.player.getPassInput(game.bot)
			} else {
				m := game.player.getInput(cells, human)
				game.apply(game.player.color, m, debug)

				for _, mv := range movelist {
					if mv.cell == m {
						game.flipDiscs(game.player.color, -mv.direction, mv.cell, debug)
					}
				}
			}
		} else if currentPlayer == white {
			movelist = nil
			cells = nil
			movelist = game.generateMoves(game.bot.color)

			fmt.Println("legal moves:")
			for _, m := range movelist {
				fmt.Printf("%v %v %v | ", color(game.bot.color, false), getCol(m.cell), getRow(m.cell))
				fmt.Printf("num filps: %v | ", m.numFlips)
				fmt.Printf("direction: %v\n", getDir(m.direction))
				cells = append(cells, m.cell)
			}

			if len(movelist) == 0 {
				fmt.Println("bot has to pass")
				if game.player.passing == false {
					game.player.passing = true
				} else {
					game.bot.passing = true
				}
			} else {
				m := game.bot.makeMoveBot(game, movelist, debug)
				if !sliceContains(m, cells) {
					fmt.Println("bot made a funny move; using rng fallback")
					m = game.bot.genRNGMove(movelist, debug)
				}

				fmt.Printf("bot generated move : %v %v %v\n", color(game.bot.color, false), getCol(m), getRow(m))
				game.apply(game.bot.color, m, debug)

				for _, mv := range movelist {
					if mv.cell == m {
						game.flipDiscs(game.bot.color, -mv.direction, mv.cell, debug)
					}
				}
			}
		}

		currentPlayer *= -1
		game.gameOver = game.isGameOver()

		turnCount++
	}

	fmt.Printf("game has ended | turns taken %v\n", turnCount)
	scores := game.calculateScoresDisc()
	printResults(scores)
}
