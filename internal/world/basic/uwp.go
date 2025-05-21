package basic

import (
	"fmt"

	"github.com/Galdoba/t5/pkg/ehex"
	"github.com/Galdoba/t5/pkg/profile"
	"github.com/Galdoba/t5/pkg/profile/prfvalue"
)

func (w *World) UWP() string {
	pr := profile.Profile(w)
	return pr.String()
}

func (w *World) Schema() (string, []*prfvalue.Value) {
	vals := []*prfvalue.Value{}
	vals = prfvalue.AddValue(vals, prfvalue.New("St", prfvalue.Ehex(w.Port)))
	vals = prfvalue.AddValue(vals, prfvalue.New("S", prfvalue.Ehex(w.Size)))
	vals = prfvalue.AddValue(vals, prfvalue.New("A", prfvalue.Ehex(w.Atmo)))
	vals = prfvalue.AddValue(vals, prfvalue.New("H", prfvalue.Ehex(w.Hydr)))
	vals = prfvalue.AddValue(vals, prfvalue.New("P", prfvalue.Ehex(w.Pops)))
	vals = prfvalue.AddValue(vals, prfvalue.New("G", prfvalue.Ehex(w.Govr)))
	vals = prfvalue.AddValue(vals, prfvalue.New("L", prfvalue.Ehex(w.Laws)))
	vals = prfvalue.AddValue(vals, prfvalue.New("-", prfvalue.Separator(false)))
	vals = prfvalue.AddValue(vals, prfvalue.New("T", prfvalue.Ehex(w.Tech)))
	return prfvalue.Schema(vals), vals
}

func (w *World) InjectValue(feed *prfvalue.Value) error {
	_, hasVals := w.Schema()
	for i, val := range hasVals {
		if i != feed.Index {
			continue
		}
		if feed.Key != val.Key {
			return fmt.Errorf("injection value key does not match (%v!=%v)", feed.Key, val.Key)
		}
		if feed.Type != val.Type {
			return fmt.Errorf("injection value type does not match (%v!=%v)", feed.Type, val.Type)
		}
		switch i {
		case 0:
			w.Port = ehex.FromString(feed.Code)
		case 1:
			w.Size = ehex.FromString(feed.Code)
		case 2:
			w.Atmo = ehex.FromString(feed.Code)
		case 3:
			w.Hydr = ehex.FromString(feed.Code)
		case 4:
			w.Pops = ehex.FromString(feed.Code)
		case 5:
			w.Govr = ehex.FromString(feed.Code)
		case 6:
			w.Laws = ehex.FromString(feed.Code)
		case 7:
			//w.Atmo = ehex.FromString(feed.Code)
		case 8:
			w.Tech = ehex.FromString(feed.Code)
		default:
			return fmt.Errorf("unexpected value index")
		}
	}
	return nil
}
