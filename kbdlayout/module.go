package kbdlayout

import (
	"time"

	"github.com/glebtv/custom_barista/kbdlayout"
	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/base"
	"github.com/soumya92/barista/base/scheduler"
)

// Module represents a clock bar module. It supports setting the click handler,
// timezone, output format, and granularity.
type Module interface {
	base.WithClickHandler
}

type module struct {
	*base.Base
	outputFunc func(time.Time) bar.Output
}

// New constructs an instance of the clock module with a default configuration.
func New() Module {
	m := &module{
		Base: base.New(),
	}
	// Default output template
	m.OnUpdate(m.update)
	return m
}

func (m *module) update() {
	now := scheduler.Now()
	m.Lock()
	var out string
	layout, err := kbdlayout.GetLayout()
	if err != nil {
		out = err.Error()
	} else {
		out = layout
	}
	m.Unlock()
	m.Output(out)
}
