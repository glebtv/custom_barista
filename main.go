// Copyright 2017 Google Inc. Apache 2.0 License
// Modifications Copyright 2018 glebtv, Apache 2.0 License
// Based on sample-bar

package main

import (
	"errors"
	"os"
	"time"

	barista "barista.run"
	"barista.run/bar"
	"barista.run/colors"
	"barista.run/modules/counter"
	"barista.run/modules/shell"
	"barista.run/outputs"
	"barista.run/pango/icons/material"
	"barista.run/pango/icons/typicons"
	"github.com/glebtv/custom_barista/batt"
	"github.com/glebtv/custom_barista/dsk"
	"github.com/glebtv/custom_barista/kbdlayout"
	"github.com/glebtv/custom_barista/load"
	"github.com/glebtv/custom_barista/ltime"
	"github.com/glebtv/custom_barista/mem"
	"github.com/glebtv/custom_barista/music"
	"github.com/glebtv/custom_barista/netm"
	"github.com/glebtv/custom_barista/temp"
	"github.com/glebtv/custom_barista/utils"
)

func main() {
	material.Load(utils.Home(".fonts/material"))
	typicons.Load(utils.Home(".fonts/typicons"))

	colors.LoadFromMap(map[string]string{
		"good":     "#6d6",
		"degraded": "#dd6",
		"bad":      "#d66",
		"dim-icon": "#777",
	})

	modules := make([]bar.Module, 0)

	modules = append(modules, kbdlayout.Get())

	//modules = append(modules, music.Get("google-play-music-desktop-player"))
	modules = append(modules, music.Get("DeaDBeeF"))

	modules = append(modules, counter.New("C:%d"))

	modules = dsk.AddTo(modules)

	modules = append(modules, load.Get())
	modules = append(modules, mem.Get())

	modules = netm.AddTo(modules)

	modules = append(modules, batt.Get())

	modules = append(modules, temp.Get())
	//modules = append(modules, weather.Get("524901"))
	//modules = append(modules, vol.Get())

	// pacin gsimplecal
	modules = append(modules, ltime.Get())

	if _, err := os.Stat("/usr/bin/rivalcfg"); errors.Is(err, os.ErrNotExist) {
		// rivalcfg does not exist
	} else {
		mouseBattery := shell.New("/usr/bin/rivalcfg", "---battery-level").
			Every(60 * time.Second).
			Output(func(count string) bar.Output {
				return outputs.Textf("%s", count)
			})
		modules = append(modules, mouseBattery)
	}

	panic(barista.Run(modules...))
}
