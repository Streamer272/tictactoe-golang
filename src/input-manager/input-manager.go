package input_manager

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TakeInput() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	return text
}
