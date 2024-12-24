package main

import (
	"fmt"

	"github.com/hw-test/makeevaas/hw02_unpack_string/unpack"
)

func main() {
	s, err := unpack.Unpack("3abc")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}
}
