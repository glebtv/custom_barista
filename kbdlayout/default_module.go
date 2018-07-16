package kbdlayout

import (
	"strings"

	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/colors"
)

func Get() bar.Module {
	return New().Output(func(i Info) bar.Output {
		out := KbdOut{}
		la := strings.ToUpper(i.Layout)
		lseg := bar.PangoSegment(la)
		if la != "US" {
			lseg.Color(colors.Scheme("bad"))
		}
		out.Seg = append(out.Seg, lseg)
		for _, mod := range i.GetMods() {
			s := bar.PangoSegment(mod)
			if mod == "CAPS" {
				s.Color(colors.Scheme("bad"))
			}
			out.Seg = append(out.Seg, s)
		}
		return out
	})
}
