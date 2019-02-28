package main
import (
	"fmt"
	"bufio"
	"os"
)
var inputReadr *bufio.Reader
var input string
var err error

func main() {
	inputReadr = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input,err = inputReadr.ReadString('\n')
	if err == nil {
		fmt.Println("The input was: %s\n",input)
	}

}