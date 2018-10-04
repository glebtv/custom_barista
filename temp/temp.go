package temp

import (
	"time"

	"barista.run/bar"
	"barista.run/colors"
	"barista.run/modules/cputemp"
	"barista.run/outputs"
	"barista.run/pango"
	"github.com/glebtv/custom_barista/utils"
	"github.com/martinlindhe/unit"
)

func Get() *cputemp.Module {
	temp := cputemp.New().
		RefreshInterval(2 * time.Second).
		Output(func(temp unit.Temperature) bar.Output {
			out := outputs.Pango(
				pango.Icon("material-build"), utils.Spacer,
				pango.Textf("%2d℃", int(temp.Celsius())),
			)
			switch {
			case temp.Celsius() > 90:
				out.Urgent(true)
			case temp.Celsius() > 70:
				out.Color(colors.Scheme("bad"))
			case temp.Celsius() > 60:
				out.Color(colors.Scheme("degraded"))
			}
			return out
		})
	return temp
}
