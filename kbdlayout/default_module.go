package kbdlayout

import (
	"strings"

	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/colors"
)

func Get() bar.Module {
	return New().OutputFunc(func(i Info) bar.Output {
		out := make(bar.Output, 0)
		la := strings.ToUpper(i.Layout)
		lseg := bar.NewSegment(la)
		if la != "US" {
			lseg.Color(colors.Scheme("bad"))
		}
		out = append(out, lseg)
		for _, mod := range i.GetMods() {
			s := bar.NewSegment(mod)
			if mod == "CAPS" {
				s.Color(colors.Scheme("bad"))
			}
			out = append(out, s)
		}
		return out
	})
}
