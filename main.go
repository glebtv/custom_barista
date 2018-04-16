// Copyright 2017 Google Inc. Apache 2.0 License
// Modifications Copyright 2018 glebtv, Apache 2.0 License
// Based on sample-bar

package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/glebtv/custom_barista/kbdlayout"
	"github.com/glebtv/custom_barista/load"
	"github.com/glebtv/custom_barista/ltime"
	"github.com/glebtv/custom_barista/mem"
	"github.com/glebtv/custom_barista/music"
	"github.com/glebtv/custom_barista/netm"
	"github.com/glebtv/custom_barista/temp"
	"github.com/glebtv/custom_barista/utils"
	"github.com/glebtv/custom_barista/weather"
	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/colors"
	"github.com/soumya92/barista/modules/battery"
	"github.com/soumya92/barista/modules/counter"
	"github.com/soumya92/barista/modules/diskio"
	"github.com/soumya92/barista/modules/diskspace"
	"github.com/soumya92/barista/outputs"
	"github.com/soumya92/barista/pango"
	"github.com/soumya92/barista/pango/icons/material"
	"github.com/soumya92/barista/pango/icons/typicons"
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
	// pacin gsimplecal

	modules := make([]bar.Module, 0)

	modules = append(modules, kbdlayout.Get())

	modules = append(modules, music.Get("google-play-music-desktop-player"))
	modules = append(modules, music.Get("DeaDBeeF"))

	modules = append(modules, counter.New("C:%d"))

	//fs := syscall.Statfs_t{}
	//err := syscall.Statfs("/", &fs)
	//if err != nil {
	//panic(err)
	//}

	modules = append(modules, diskspace.New("/").OutputTemplate(outputs.TextTemplate(`FREE / {{.Free.Gigabytes | printf "%.2f"}} GB`)))

	path, err := exec.LookPath("findmnt")
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(path, "/")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	} else {
		parts := strings.Split(string(out), "\n")
		info := strings.Fields(parts[1])
		dn := strings.Split(info[1], "/")
		name := dn[len(dn)-1]

		diskio.RefreshInterval(2 * time.Second)

		sda := diskio.Disk(name).
			OutputFunc(func(io diskio.IO) bar.Output {
				//spew.Dump(io)
				return outputs.Pango(
					pango.Textf("io"),
					pango.Textf("%9s", outputs.IByterate(io.Input)),
					utils.Spacer,
					pango.Textf("%9s", outputs.IByterate(io.Output)),
				)
			})

		modules = append(modules, sda)
	}

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
