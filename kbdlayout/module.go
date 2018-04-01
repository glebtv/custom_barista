package kbdlayout

import (
	"strings"

	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/base"
)

const NUM_LOCK = 16
const CAPS_LOCK = 2

type Info struct {
	// Layout code
	Layout string
	// Mods - modifier map
	Mods uint8
}

func (i Info) GetMods() []string {
	ret := make([]string, 0)
	if i.Mods&NUM_LOCK == NUM_LOCK {
		ret = append(ret, "NUM")
	}
	if i.Mods&CAPS_LOCK == CAPS_LOCK {
		ret = append(ret, "CAPS")
	}
	//log.Println("getmods", mods, ret)
	return ret
}

type Module struct {
	*base.Base
	outputFunc func(Info) bar.Output
}

var DefaultOutputFunc = func(i Info) bar.Output {
	out := make(bar.Output, 0)
	lseg := bar.NewSegment(strings.ToUpper(i.Layout))
	out = append(out, lseg)
	for _, mod := range i.GetMods() {
		out = append(out, bar.NewSegment(mod))
	}
	return out
}

// New constructs an instance of the clock module with a default configuration.
func New() *Module {
	m := &Module{
		Base:       base.New(),
		outputFunc: DefaultOutputFunc,
	}
	// Default output template
	m.OnUpdate(m.update)

	Subscribe(func(layout string, mods uint8) {
		i := Info{Layout: layout, Mods: mods}
		m.Output(m.outputFunc(i))
	})

	return m
}

func (m *Module) OutputFunc(outputFunc func(Info) bar.Output) *Module {
	m.Lock()
	defer m.UnlockAndUpdate()
	m.outputFunc = outputFunc
	return m
}

func (m *Module) update() {
	m.Lock()
	layout, mods, err := GetLayout()
	if err != nil {
		layout = err.Error()
		mods = 0
	}
	m.Unlock()
	i := Info{Layout: layout, Mods: mods}
	m.Output(m.outputFunc(i))
}

func (m *Module) Click(e bar.Event) {
	if e.Button == bar.ButtonLeft {
		SwitchToNext()
		m.update()
	}
	m.Base.Click(e)
}
