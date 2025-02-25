package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rodaine/numwords"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	textInput, _ := reader.ReadString('\n')

	fmt.Println(numwords.ParseString(textInput))
}
