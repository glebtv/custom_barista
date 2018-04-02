package mem

import (
	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/colors"
	"github.com/soumya92/barista/modules/meminfo"
	"github.com/soumya92/barista/outputs"
	"github.com/soumya92/barista/pango/icons/material"
)

func Get() bar.Module {
	return meminfo.New().OutputFunc(func(m meminfo.Info) bar.Output {
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
}
