package main

import (
	"fmt"
	"github.com/cloudfoundry-community/goutils/tree"
)

func main() {
	t := tree.New("a",
		tree.New("b"),
		tree.New("c"),
	)

	fmt.Printf("%s\n", t.Draw())
}
