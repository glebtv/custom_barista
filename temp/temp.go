package temp

import (
	"time"

	"github.com/glebtv/custom_barista/utils"
	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/colors"
	"github.com/soumya92/barista/modules/cputemp"
	"github.com/soumya92/barista/outputs"
	"github.com/soumya92/barista/pango"
	"github.com/soumya92/barista/pango/icons/material"
)

func Get() cputemp.Module {
	temp := cputemp.DefaultZone().
		RefreshInterval(2 * time.Second).
		UrgentWhen(func(temp cputemp.Temperature) bool {
			return temp.C() > 90
		}).
		OutputColor(func(temp cputemp.Temperature) bar.Color {
			switch {
			case temp.C() > 70:
				return colors.Scheme("bad")
			case temp.C() > 60:
				return colors.Scheme("degraded")
			default:
				return colors.Empty()
			}
		}).
		OutputFunc(func(temp cputemp.Temperature) bar.Output {
			return outputs.Pango(
				material.Icon("build"), utils.Spacer,
				pango.Textf("%2d℃", temp.C()),
			)
		})
	return temp
}
