package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func ReadInput() int {
	scanner.Scan()
	text := scanner.Text()
	n, err := strconv.Atoi(text)

	for err != nil {
		fmt.Println(err.Error())
		scanner.Scan()
		text = scanner.Text()
		n, err = strconv.Atoi(text)
	}

	return n
}
