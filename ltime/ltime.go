package ltime

import (
	"time"

	"barista.run/bar"
	"barista.run/base/click"
	"barista.run/colors"
	"barista.run/modules/clock"
	"barista.run/outputs"
	"barista.run/pango"
)

func Get() bar.Module {
	time := clock.Local().Output(time.Second, func(now time.Time) bar.Output {
		return outputs.Pango(
			pango.Icon("material-today").Color(colors.Scheme("dim-icon")),
			now.Format("Mon 2006-01-02 "),
			pango.Icon("material-access-time").Color(colors.Scheme("dim-icon")),
			now.Format("15:04:05"),
		).OnClick(click.RunLeft("gsimplecal"))
	})
	return time
}
