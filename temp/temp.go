package temp

import (
	"image/color"
	"time"

	"github.com/glebtv/custom_barista/utils"
	"github.com/martinlindhe/unit"
	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/colors"
	"github.com/soumya92/barista/modules/cputemp"
	"github.com/soumya92/barista/outputs"
	"github.com/soumya92/barista/pango"
)

func Get() *cputemp.Module {
	temp := cputemp.DefaultZone().
		RefreshInterval(2 * time.Second).
		UrgentWhen(func(temp unit.Temperature) bool {
			return temp.Celsius() > 90
		}).
		OutputColor(func(temp unit.Temperature) color.Color {
			switch {
			case temp.Celsius() > 70:
				return colors.Scheme("bad")
			case temp.Celsius() > 60:
				return colors.Scheme("degraded")
			default:
				return nil
			}
		}).
		OutputFunc(func(temp unit.Temperature) bar.Output {
			return outputs.Pango(
				pango.Icon("material-build"), utils.Spacer,
				pango.Textf("%2d℃", int(temp.Celsius())),
			)
		})
	return temp
}
