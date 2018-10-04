package main

import (
	"github.com/glebtv/custom_barista/kbdlayout"
	"barista.run/bar"
)

func main() {
	//layout := kbdlayout.New()
	layout := kbdlayout.Get()

	panic(bar.Run(
		layout,
	))
}
