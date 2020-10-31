package main

// Move A move struct to hold info about a move
type Move struct {
	cell      int
	numFlips  int
	direction int // need to think about making this a slice and have multiple directions
}

// GetWeight Gets the weight of a move's cell
func (m Move) GetWeight() int {
	return weights[m.cell]
}

func (b Board) getLegalMoves(index, color, dir int) Move {
	flips := 0
	i := index
	d := dir
	m := Move{cell: 0, numFlips: 0, direction: 0}
	flag := false

Check:
	for index >= 0 && index < int(boardSize) && !flag {
		flag = checkDir(i, d)
		i += d

		if i >= 0 && i < int(boardSize) {
			if b.board[uint(i)] != -color {
				break Check
			} else {
				flips++
			}
		} else {
			flips = 0
			break Check
		}
	}

	if i >= 0 && i < int(boardSize) {
		if b.board[uint(i)] == 0 && flips != 0 {
			m.cell = i
			m.numFlips = flips
			m.direction = dir
		}
	}

	return m
}

func (b Board) generateMoves(color int) []Move {
	var moveset []Move

	for i, val := range b.board {
		if val == color {
			for _, d := range directions {
				temp := b.getLegalMoves(i, color, d)

				if temp.numFlips != 0 && !checkBorderMove(temp) {
					moveset = append(moveset, temp)
				}
			}
		}
	}

	return moveset
}

func checkBorderMove(m Move) bool {
	switch {
	case contains(m.cell, leftBorder):
		if m.direction == -west || m.direction == -northWest || m.direction == -southWest {
			return true
		}
	case contains(m.cell, rightBorder):
		if m.direction == -east || m.direction == -northEast || m.direction == -southEast {
			return true
		}
	}
	return false
}

func checkDir(index, direction int) bool {
	switch {
	case direction == east && contains(index, rightBorder):
		return true
	case direction == west && contains(index, leftBorder):
		return true
	case direction == northEast && (contains(index, rightBorder) || contains(index, topBorder)):
		return true
	case direction == northWest && (contains(index, leftBorder) || contains(index, topBorder)):
		return true
	case direction == southEast && (contains(index, rightBorder) || contains(index, bottomBorder)):
		return true
	case direction == southWest && (contains(index, leftBorder) || contains(index, bottomBorder)):
		return true
	default:
		return false
	}
}
