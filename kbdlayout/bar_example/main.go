package main

import (
	"github.com/glebtv/custom_barista/kbdlayout"
	"github.com/soumya92/barista/bar"
)

func main() {
	//layout := kbdlayout.New()
	layout := kbdlayout.Get()

	panic(bar.Run(
		layout,
	))
}
