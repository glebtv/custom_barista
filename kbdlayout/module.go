package kbdlayout

import (
	"strings"

	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/base"
)

type module struct {
	*base.Base
}

// New constructs an instance of the clock module with a default configuration.
func New() bar.Module {
	m := &module{
		Base: base.New(),
	}
	// Default output template
	m.OnUpdate(m.update)

	Subscribe(func(layout string, mods uint8) {
		m.Send(layout, mods)
	})

	return m
}

func (m *module) update() {
	m.Lock()
	layout, mods, err := GetLayout()
	if err != nil {
		layout = err.Error()
		mods = 0
	}
	m.Unlock()
	m.Send(layout, mods)
}

func (m *module) Send(layout string, mods uint8) {
	out := make(bar.Output, 0)
	lseg := bar.NewSegment(strings.ToUpper(layout))
	out = append(out, lseg)
	for _, mod := range GetMods(mods) {
		out = append(out, bar.NewSegment(mod))
	}
	m.Output(out)
}

func (m *module) Click(e bar.Event) {
	if e.Button == bar.ButtonLeft {
		SwitchToNext()
		m.update()
	}
	m.Base.Click(e)
}
