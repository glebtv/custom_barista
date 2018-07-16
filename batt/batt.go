package batt

import (
	"time"

	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/colors"
	"github.com/soumya92/barista/modules/battery"
	"github.com/soumya92/barista/outputs"
)

func Get() bar.Module {
	statusName := map[battery.Status]string{
		battery.Charging:    "CHR",
		battery.Discharging: "BAT",
		battery.NotCharging: "NOT",
		battery.Unknown:     "UNK",
	}

	return battery.All().Output(func(b battery.Info) bar.Output {
		if b.Status == battery.Disconnected {
			return nil
		}
		if b.Status == battery.Full {
			return outputs.Text("FULL")
		}
		out := outputs.Textf("%s %d%% %s",
			statusName[b.Status],
			b.RemainingPct(),
			b.RemainingTime())
		if b.Discharging() {
			if b.RemainingPct() < 20 || b.RemainingTime() < 30*time.Minute {
				out.Color(colors.Scheme("bad"))
			}
		}
		return out
	})
}
