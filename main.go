package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 4 {
		i, _ := strconv.Atoi(os.Args[1])
		j := os.Args[2]
		k, _ := strconv.Atoi(os.Args[3])

		switch j {
		case "+":
			fmt.Print(i + k)
			return
		case "-":
			fmt.Print(i - k)
			return
		case "*":
			fmt.Print(i * k)
			return
		}
	}
	fmt.Print("使い方:<a> <+ or - or *> <b>")
}
