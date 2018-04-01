// Copyright 2017 Google Inc. Apache 2.0 License
// Modifications Copyright 2018 glebtv, Apache 2.0 License
// Based on sample-bar

package main

import (
	"os/exec"
	"os/user"
	"path/filepath"
	"time"

	"github.com/glebtv/custom_barista/kbdlayout"
	"github.com/glebtv/custom_barista/temp"
	"github.com/glebtv/custom_barista/weather"
	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/colors"
	"github.com/soumya92/barista/modules/clock"
	"github.com/soumya92/barista/modules/group"
	"github.com/soumya92/barista/modules/meminfo"
	"github.com/soumya92/barista/modules/netspeed"
	"github.com/soumya92/barista/modules/sysinfo"
	"github.com/soumya92/barista/modules/wlan"
	"github.com/soumya92/barista/outputs"
	"github.com/soumya92/barista/pango"
	"github.com/soumya92/barista/pango/icons/fontawesome"
	"github.com/soumya92/barista/pango/icons/material"
)

var spacer = pango.Span(" ", pango.XXSmall)

func home(path string) string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return filepath.Join(usr.HomeDir, path)
}

func main() {
	// git clone git@github.com:google/material-design-icons.git ~/material-design-icons
	material.Load(home("material-design-icons"))

	colors.LoadFromMap(map[string]string{
		"good":     "#6d6",
		"degraded": "#dd6",
		"bad":      "#d66",
		"dim-icon": "#777",
	})
	// pacin gsimplecal

	localtime := clock.New().OutputFunc(func(now time.Time) bar.Output {
		return outputs.Pango(
			material.Icon("today", colors.Scheme("dim-icon")),
			now.Format("2006-01-02 "),
			material.Icon("access-time", colors.Scheme("dim-icon")),
			now.Format("15:04:05"),
		)
	}).OnClick(func(e bar.Event) {
		if e.Button == bar.ButtonLeft {
			exec.Command("gsimplecal").Run()
		}
	})

	loadAvg := sysinfo.New().OutputFunc(func(s sysinfo.Info) bar.Output {
		out := outputs.Textf("%0.2f %0.2f", s.Loads[0], s.Loads[2])
		// Load averages are unusually high for a few minutes after boot.
		if s.Uptime < 10*time.Minute {
			// so don't add colours until 10 minutes after system start.
			return out
		}
		switch {
		case s.Loads[0] > 128, s.Loads[2] > 64:
			out.Urgent(true)
		case s.Loads[0] > 64, s.Loads[2] > 32:
			out.Color(colors.Scheme("bad"))
		case s.Loads[0] > 32, s.Loads[2] > 16:
			out.Color(colors.Scheme("degraded"))
		}
		return out
	})

	freeMem := meminfo.New().OutputFunc(func(m meminfo.Info) bar.Output {
		out := outputs.Pango(material.Icon("memory"), m.Available().IEC())
		freeGigs := m.Available().In("GiB")
		switch {
		case freeGigs < 0.5:
			out.Urgent(true)
		case freeGigs < 1:
			out.Color(colors.Scheme("bad"))
		case freeGigs < 2:
			out.Color(colors.Scheme("degraded"))
		case freeGigs > 12:
			out.Color(colors.Scheme("good"))
		}
		return out
	})

	net := netspeed.New("enp4s0").
		RefreshInterval(2 * time.Second).
		OutputFunc(func(s netspeed.Speeds) bar.Output {
			return outputs.Pango(
				fontawesome.Icon("file_upload"), spacer, pango.Textf("%5s", s.Tx.SI()),
				pango.Span(" ", pango.Small),
				fontawesome.Icon("file_download"), spacer, pango.Textf("%5s", s.Rx.SI()),
			)
		})

	wlan := wlan.New("wlp3s0")

	layout := kbdlayout.New()

	g := group.Collapsing()

	panic(bar.Run(
		layout,
		g.Add(net),
		g.Add(wlan),
		g.Add(temp.Get()),
		g.Add(freeMem),
		g.Add(loadAvg),
		weather.Get(),
		localtime,
	))
}
