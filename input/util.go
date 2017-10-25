package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func ReadInput() (int, error) {
	scanner.Scan()
	text := scanner.Text()
	return strconv.Atoi(text)
}

func ReadYesNo() bool {
	fmt.Printf("(Y)es or (N)o:")
	scanner.Scan()
	text := scanner.Text()
	return (text == "Y" || text == "y")
}
