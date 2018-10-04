package mem

import (
	"barista.run/bar"
	"barista.run/colors"
	"barista.run/modules/meminfo"
	"barista.run/outputs"
	"barista.run/pango"
)

// Get Create a module
func Get() bar.Module {
	return meminfo.New().Output(func(m meminfo.Info) bar.Output {
		out := outputs.Pango(pango.Icon("material-memory"), outputs.IBytesize(m.Available()))
		freeGigs := m.Available().Gigabytes()
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
