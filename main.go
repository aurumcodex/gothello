package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

func main() {
	// need to get and parse cli arguments
	cores := runtime.NumCPU()
	board := Board{}

	b := board

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	fmt.Printf("you entered: %v\n", text)

	for i, c := range []rune(text) {
		fmt.Printf("%v: %c\n", i, c)
	}

	fmt.Println("derp")
	fmt.Printf("you have: %v cores available\n", cores)
	fmt.Println(board)
	fmt.Println(b)
	board.board[33] = 11
	fmt.Println(board)

	fmt.Printf("%c\n", disc)
	fmt.Println(maxInt)
	fmt.Println(minInt)
}
