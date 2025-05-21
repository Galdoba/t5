package wrld

import (
	"fmt"

	"github.com/Galdoba/t5/pkg/ehex"
	"github.com/Galdoba/t5/pkg/profile/prfvalue"
)

type World struct {
	S ehex.Ehex
	A ehex.Ehex
	H ehex.Ehex
}

func (w *World) Schema() (string, []*prfvalue.Value) {
	vals := []*prfvalue.Value{}
	vals = prfvalue.AddValue(vals, prfvalue.New("S", prfvalue.Ehex(w.S)))
	vals = prfvalue.AddValue(vals, prfvalue.New("A", prfvalue.Ehex(w.A)))
	vals = prfvalue.AddValue(vals, prfvalue.New("H", prfvalue.Ehex(w.H)))
	return prfvalue.Schema(vals), vals
}

func (w *World) String() string {
	return fmt.Sprintf("-%v-%v-%v-", w.S.Code(), w.A.Code(), w.H.Code())
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
			w.S = ehex.FromString(feed.Code)
		case 1:
			w.A = ehex.FromString(feed.Code)
		case 2:
			w.H = ehex.FromString(feed.Code)
		default:
			return fmt.Errorf("unexpected value index")
		}
	}
	return nil
}
