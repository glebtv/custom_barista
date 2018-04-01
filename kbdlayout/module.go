package kbdlayout

import (
	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/base"
	"github.com/soumya92/barista/outputs"
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
	return m
}

func (m *module) update() {
	m.Lock()
	var out string
	layout, err := GetLayout()
	if err != nil {
		out = err.Error()
	} else {
		out = layout
	}
	m.Unlock()
	m.Output(outputs.Text(out))

	Subscribe(func(layout string) {
		m.Output(outputs.Text(layout))
	})
}

func (m *module) Click(e bar.Event) {
	if e.Button == bar.ButtonLeft {
		SwitchToNext()
		m.update()
	}
	m.Base.Click(e)
}
