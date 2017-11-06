package io

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

type Console struct {
}

func (c *Console) ReadBool() bool {
	fmt.Printf("(Y)es or (N)o:")
	scanner.Scan()
	text := scanner.Text()
	return (text == "Y" || text == "y")
}

func (c *Console) ReadInt() int {
	for {
		scanner.Scan()
		text := scanner.Text()
		n, err := strconv.Atoi(text)
		if err == nil {
			return n
		}
	}
}

func (c *Console) Write(msg interface{}) {
	fmt.Print(msg)
}
