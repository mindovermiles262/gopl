// Echo1 prints it's command line arguments
package main

import (
  "fmt"
  "os"
)

func main() {
  var st, sep string
  for i := 1; i < len(os.Args); i++ {
    st += sep + os.Args[i]
    sep = " "
  }
  fmt.Println(st)
}

