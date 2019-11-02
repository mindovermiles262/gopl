// Echo2 prints it's command line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	st, sep := "", ""
	for _, arg := range os.Args[1:] {
		st += sep + arg
		sep = " "
	}
	fmt.Println(st)
}
