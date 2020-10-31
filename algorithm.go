package main

import (
	"fmt"
	"math"
)

//implement search algorithms here

func (b Board) alphaBeta(alpha, beta float64, player, depth /*turnCount*/ int, maxing, debug bool) int {
	// get fresh copies of alpha and beta (maybe)

	moveCount := len(b.generateMoves(player))
	bestMove := -1

	if debug {
		fmt.Println("moves available:", moveCount, "| depth =", depth)
	}

	if depth == maxDepth {
		if debug {
			fmt.Println("hit max depth (15)")
		}

		scores := b.calculateScores()
		bestMove = scores.Score

		if debug {
			b.print(emptyMoves)
		}
	} else if depth < maxDepth {
		if maxing {
			bestMove = minInt
			moveset := b.generateMoves(player)

		Max:
			for _, m := range moveset {
				if debug {
					fmt.Println("Legal Cell =", m.cell)
				}

				temp := b
				temp.apply(player, m.cell, debug)
				temp.flipDiscs(player, -1*m.direction, m.cell, debug)

				val := temp.alphaBeta(alpha, beta, -player, depth+1, !maxing, debug)

				bestMove = max(bestMove, val)
				alpha = math.Max(alpha, float64(bestMove))

				if alpha >= beta {
					break Max
				}
			}
		} else if !maxing {
			bestMove = maxInt
			moveset := b.generateMoves(player)

		Min:
			for _, m := range moveset {
				if debug {
					fmt.Println("Legal Cell =", m.cell)
				}

				temp := b
				temp.apply(player, m.cell, debug)
				temp.flipDiscs(player, -1*m.direction, m.cell, debug)

				val := temp.alphaBeta(alpha, beta, -player, depth+1, !maxing, debug)

				bestMove = min(bestMove, val)
				beta = math.Min(beta, float64(bestMove))

				if alpha >= beta {
					break Min
				}
			}
		}
	}

	return bestMove
} // end alphaBeta()

// Negamax is a fuction
func (b Board) negamax(alpha, beta float64, player, depth int, debug bool) int {
	moveset := b.generateMoves(player)
	moveCount := len(moveset)
	bestMove := minInt

	if debug {
		fmt.Println("moves available:", moveCount, "| depth =", depth)
	}

	if depth == 0 {
		return player * b.calculateScores().Score
	}

Cycle:
	for _, m := range moveset {
		if debug {
			fmt.Println("Legal Cell =", m.cell)
		}

		temp := b
		temp.apply(player, m.cell, debug)
		temp.flipDiscs(player, -1*m.direction, m.cell, debug)

		bestMove = max(bestMove, -temp.negamax(-beta, -alpha, -player, depth-1, debug))
		alpha = math.Max(alpha, float64(bestMove))

		if alpha >= beta {
			break Cycle
		}
	}

	return bestMove
}
