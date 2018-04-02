// Copyright 2017 Google Inc. Apache 2.0 License
// Modifications Copyright 2018 glebtv, Apache 2.0 License
// Based on sample-bar

package main

import (
	"os"

	"github.com/glebtv/custom_barista/kbdlayout"
	"github.com/glebtv/custom_barista/load"
	"github.com/glebtv/custom_barista/ltime"
	"github.com/glebtv/custom_barista/mem"
	"github.com/glebtv/custom_barista/netm"
	"github.com/glebtv/custom_barista/temp"
	"github.com/glebtv/custom_barista/utils"
	"github.com/glebtv/custom_barista/weather"
	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/colors"
	"github.com/soumya92/barista/modules/battery"
	"github.com/soumya92/barista/pango/icons/material"
	"github.com/soumya92/barista/pango/icons/typicons"
)

func main() {
	material.Load(utils.Home("material-design-icons"))
	typicons.Load(utils.Home(".fonts/typicons"))

	colors.LoadFromMap(map[string]string{
		"good":     "#6d6",
		"degraded": "#dd6",
		"bad":      "#d66",
		"dim-icon": "#777",
	})
	// pacin gsimplecal

	modules := make([]bar.Module, 0)

	modules = append(modules, kbdlayout.Get())

	modules = append(modules, load.Get())

	modules = append(modules, mem.Get())

	modules = netm.AddTo(modules)

	if _, err := os.Stat("/sys/class/power_supply/BAT0"); err == nil {
		modules = append(modules, battery.New("BAT0"))
	}

	modules = append(modules, temp.Get())
	modules = append(modules, weather.Get())

	modules = append(modules, ltime.Get())

	panic(bar.Run(modules...))
}
