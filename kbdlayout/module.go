package kbdlayout

import (
	"strings"
	"sync"

	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/base"
)

const NUM_LOCK = 16
const CAPS_LOCK = 2

var lock sync.Mutex

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
	bar.Module
	bar.Sink
	base.SimpleClickHandler
	output func(Info) bar.Output
}

type KbdOut struct {
	Seg []*bar.Segment
}

func (k KbdOut) Segments() []*bar.Segment {
	return k.Seg
}

var DefaultOutput = func(i Info) bar.Output {
	out := KbdOut{}
	lseg := bar.TextSegment(strings.ToUpper(i.Layout))
	out.Seg = append(out.Seg, lseg)
	for _, mod := range i.GetMods() {
		out.Seg = append(out.Seg, bar.TextSegment(mod))
	}
	return bar.Output(out)
}

func (m *Module) Stream(s bar.Sink) {
	forever := make(chan struct{})
	m.Sink = s
	<-forever
}

// New constructs an instance of the clock module with a default configuration.
func New() *Module {
	m := &Module{
		output: DefaultOutput,
	}
	// Default output template

	Subscribe(func(layout string, mods uint8) {
		i := Info{Layout: layout, Mods: mods}
		m.Sink.Output(m.output(i))
	})

	return m
}

func (m *Module) Output(output func(Info) bar.Output) *Module {
	m.output = output
	return m
}

func (m *Module) update() {
	layout, mods, err := GetLayout()
	if err != nil {
		layout = err.Error()
		mods = 0
	}
	i := Info{Layout: layout, Mods: mods}
	m.Sink.Output(m.output(i))
}

func (m *Module) Click(e bar.Event) {
	if e.Button == bar.ButtonLeft {
		SwitchToNext()
		m.update()
	}
}
