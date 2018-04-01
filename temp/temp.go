package temp

import (
	"os/user"
	"path/filepath"
	"time"

	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/colors"
	"github.com/soumya92/barista/modules/cputemp"
	"github.com/soumya92/barista/outputs"
	"github.com/soumya92/barista/pango"
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

func Get() cputemp.Module {
	material.Load(home("material-design-icons"))

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
				material.Icon("build"), spacer,
				pango.Textf("%2d℃", temp.C()),
			)
		})
	return temp
}
