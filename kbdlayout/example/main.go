package main

import (
	"fmt"

	"github.com/glebtv/custom_barista/kbdlayout"
)

func main() {
	layout, err := kbdlayout.GetLayout()
	if err != nil {
		panic(err)
	}
	fmt.Println("layout:", layout)
}
