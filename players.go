package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Player struct {
	color    int
	numDiscs int
	score    int
	human    bool
	passing  bool
}

func (p Player) init(c int, h bool) Player {
	return Player{
		color: c,
		human: h,
	}
}

// [===== human input logic section =====]

func (p Player) getInput(cells []int, human bool) int {
	empty := func(l []int) bool {
		return len(l) == 0
	}

	move := maxInt
	row := 0
	col := 0

	reader := bufio.NewReader(os.Stdin)

	if !empty(cells) {
		fmt.Println("player has vaild moves available.")
	}

	fmt.Print("Enter a move (color, column, row): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to read user input")
		os.Exit(2)
	}
	input = strings.Replace(input, "\n", "", -1)

	chars := strings.Split(input, " ")

	// need to figure out if it's possible to simplify this logic even more
	if (chars[0] == "B" && p.color == black) && !empty(cells) && len(chars) > 1 {
		row = rows[chars[2]]
		col = columns[chars[1]]
		move = (row * 8) + col

		if !sliceContains(move, cells) {
			switch human {
			case true:
				fmt.Println("since a human is playing, re-enter move")
				p.getInput(cells, human)
			case false:
				fmt.Fprintln(os.Stderr, "invalid move entered")
				os.Exit(1)
			}
		}
	} else if (chars[0] == "W" && p.color == white) && !empty(cells) && len(chars) > 1 {
		row = rows[chars[2]]
		col = columns[chars[1]]
		move = (row * 8) + col

		if !sliceContains(move, cells) {
			switch human {
			case true:
				fmt.Println("since a human is playing, re-enter move")
				p.getInput(cells, human)
			case false:
				fmt.Fprintln(os.Stderr, "invalid move entered")
				os.Exit(1)
			}
		}
	} else {
		// fmt.Println
		fmt.Fprintln(os.Stderr, "invalid move entered")
		os.Exit(1)
	}

	fmt.Printf("player %v made move at cell: %v\n", color(p.color, fullColor), move)
	return move
}

func (p Player) getPassInput(opponent Player) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to read user input")
		os.Exit(2)
	}
	input = strings.Replace(input, "\n", "", -1)

	// need to clean up boilerplate here
	switch input {
	case "b":
		p.handleInputBlack(opponent)
	case "B":
		p.handleInputBlack(opponent)
	case "w":
		p.handleInputWhite(opponent)
	case "W":
		p.handleInputWhite(opponent)
	default:
		if p.human {
			fmt.Println("invalid option found. please re-enter")
			p.getPassInput(opponent)
		}
	}
}

func (p Player) handleInputBlack(opponent Player) {
	if p.color != black && p.human {
		fmt.Println("you have no valid moves and need to pass. re-enter input")
		p.getPassInput(opponent)
	}
	switch p.passing {
	case true:
		opponent.passing = true
	case false:
		p.passing = true
	}
}

func (p Player) handleInputWhite(opponent Player) {
	if p.color != white && p.human {
		fmt.Println("you have no valid moves and need to pass. re-enter input")
		p.getPassInput(opponent)
	}
	switch p.passing {
	case true:
		opponent.passing = true
	case false:
		p.passing = true
	}
}

// [===== bot logic section =====]
func (p Player) makeMoveBot(b Board, random, debug bool) int {
	// TODO: add logic
	moveset := b.generateMoves(p.color)
	// cells := getCells(moveset)

	move := -1
	// depth := 0
	// maxing := true
	// alpha := float64(math.MinInt64)
	// beta := float64(math.MaxInt64)

	if random {
		move = p.genRNGMove(moveset, debug)
	} else {

	}

	return move // for now
}

func (p Player) genRNGMove(moveset []Move, debug bool) int {
	rand.Seed(time.Now().UnixNano())

	cells := getCells(moveset)
	move := rand.Intn(boardSize)

	if debug {
		fmt.Println("unsorted cell list:", cells)
	}

	for !sliceContains(move, cells) {
		move = rand.Intn(boardSize)
	}

	return move
}

// func genFallbackMove(cells []int) int {
// 	rand.Seed(time.Now().UnixNano())
// 	move := rand.Intn(boardSize)

// 	for !sliceContains(move, cells) {
// 		move = rand.Intn(boardSize)
// 	}

// 	return move
// }
