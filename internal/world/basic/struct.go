package basic

import (
	"fmt"

	"github.com/Galdoba/t5/pkg/ehex"
)

type World struct {
	MainWorld          bool
	HabitableZoneOrbit int
	Template           string
	Port               ehex.Ehex
	Size               ehex.Ehex
	Atmo               ehex.Ehex
	Hydr               ehex.Ehex
	Pops               ehex.Ehex
	Govr               ehex.Ehex
	Laws               ehex.Ehex
	Tech               ehex.Ehex
}

func New(options ...WorldGenerationOption) *World {
	w := World{}
	for _, set := range options {
		set(&w)
	}
	return &w
}

type WorldGenerationOption func(*World)

func MainWorld(mw bool) WorldGenerationOption {
	return func(w *World) {
		w.MainWorld = mw
	}
}

func HabitableZone(hz int) WorldGenerationOption {
	return func(w *World) {
		w.HabitableZoneOrbit = hz
	}
}

func Template(tmpl string) WorldGenerationOption {
	return func(w *World) {
		w.Template = tmpl
	}
}

////////////////////////

func (w *World) String() string {
	if w.Size == nil {
		w.Size = ehex.FromString("X")
	}
	if w.Atmo == nil {
		w.Atmo = ehex.FromString("X")
	}
	if w.Hydr == nil {
		w.Hydr = ehex.FromString("X")
	}
	if w.Pops == nil {
		w.Pops = ehex.FromString("X")
	}
	if w.Govr == nil {
		w.Govr = ehex.FromString("X")
	}
	if w.Laws == nil {
		w.Laws = ehex.FromString("X")
	}
	if w.Port == nil {
		w.Port = ehex.FromString("X")
	}
	if w.Tech == nil {
		w.Tech = ehex.FromString("X")
	}
	return fmt.Sprintf("%v\t%v\t%v", w.Template, w.UWP(), w.HabitableZoneOrbit)
}
