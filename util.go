package main

import (
	"bufio"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func ReadInput() int {
	scanner.Scan()
	text := scanner.Text()
	n, err := strconv.Atoi(text)
	for err != nil {
		scanner.Scan()
		text = scanner.Text()
		n, err = strconv.Atoi(text)
	}

	return n
}
