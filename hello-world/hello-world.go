package main

import (
	"fmt"
)
// go build builds this just fine, i.e
// the compile does not care about the
// formating. But `go fmt hello-world.go`
// wants that it looks different
func main() {fmt.Println("Hello world!")
            }
