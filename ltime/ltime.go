package ltime

import (
	"os/exec"
	"time"

	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/colors"
	"github.com/soumya92/barista/modules/clock"
	"github.com/soumya92/barista/outputs"
	"github.com/soumya92/barista/pango/icons/material"
)

func Get() bar.Module {
	return clock.New().OutputFunc(func(now time.Time) bar.Output {
		return outputs.Pango(
			material.Icon("today", colors.Scheme("dim-icon")),
			now.Format("Mon 2006-01-02 "),
			material.Icon("access-time", colors.Scheme("dim-icon")),
			now.Format("15:04:05"),
		)
	}).OnClick(func(e bar.Event) {
		if e.Button == bar.ButtonLeft {
			exec.Command("gsimplecal").Run()
		}
	})
}
