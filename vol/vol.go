package vol

import (
	"github.com/glebtv/custom_barista/utils"
	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/colors"
	"github.com/soumya92/barista/modules/volume"
	"github.com/soumya92/barista/outputs"
	"github.com/soumya92/barista/pango"
)

func Get() *volume.Module {
	return volume.DefaultMixer().Output(func(v volume.Volume) bar.Output {
		if v.Mute {
			return outputs.
				Pango(pango.Icon("ion-volume-off"), "MUT").
				Color(colors.Scheme("degraded"))
		}
		iconName := "mute"
		pct := v.Pct()
		if pct > 66 {
			iconName = "high"
		} else if pct > 33 {
			iconName = "low"
		}
		return outputs.Pango(
			pango.Icon("ion-volume-"+iconName),
			utils.Spacer,
			pango.Textf("%2d%%", pct),
		)
	})
}
