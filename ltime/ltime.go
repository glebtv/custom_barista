package ltime

import (
	"os/exec"
	"time"

	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/colors"
	"github.com/soumya92/barista/modules/clock"
	"github.com/soumya92/barista/outputs"
	"github.com/soumya92/barista/pango"
)

func Get() bar.Module {
	time := clock.Local().Output(time.Second, func(now time.Time) bar.Output {
		return outputs.Pango(
			pango.Icon("material-today").Color(colors.Scheme("dim-icon")),
			now.Format("Mon 2006-01-02 "),
			pango.Icon("material-access-time").Color(colors.Scheme("dim-icon")),
			now.Format("15:04:05"),
		)
	})
	time.OnClick(func(e bar.Event) {
		if e.Button == bar.ButtonLeft {
			exec.Command("gsimplecal").Run()
		}
	})
	return time
}
