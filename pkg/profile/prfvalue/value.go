package prfvalue

import (
	"fmt"

	"github.com/Galdoba/t5/pkg/ehex"
)

type Value struct {
	Key              string  `json:"key,omitempty"`
	Index            int     `json:"index,omitempty"`
	Type             string  `json:"type,omitempty"`
	Code             string  `json:"Code,omitempty"`
	ValueInt         int     `json:"ValueInt,omitempty"`
	ValueFloat64     float64 `json:"ValueFloat64,omitempty"`
	ValueBool        bool    `json:"ValueBool,omitempty"`
	ValueDescription string  `json:"ValueDescription,omitempty"`
	Omitable         bool    `json:"can be omited,omitempty"`
	isSet            bool
}

func New(key string, feed ...ValueFeed) *Value {
	val := Value{}
	val.Key = key
	for _, set := range feed {
		set(&val)
	}
	if !val.isSet {
		// return nil
	}
	return &val
}

func AddValue(vals []*Value, val *Value) []*Value {
	val.Index = len(vals)
	vals = append(vals, val)
	return vals
}

func Schema(vals []*Value) string {
	sc := ""
	for _, v := range vals {
		sc += v.Key
	}
	return sc
}

type ValueFeed func(*Value)

func Ehex(eh ehex.Ehex) ValueFeed {
	return func(v *Value) {
		if v.isSet {
			return
		}
		v.Type = "ehex"
		v.Code = eh.Code()
		v.ValueInt = eh.Value()
		v.isSet = true
	}
}

func Int(i int, precision ...int) ValueFeed {
	return func(v *Value) {
		if v.isSet {
			return
		}
		prec := 0
		for _, p := range precision {
			prec = p
		}
		v.Type = "int"
		v.Code = formatInt(i, prec)
		v.ValueInt = i
		v.isSet = true
	}
}

func formatInt(i int, prec int) string {
	return fmt.Sprintf("%v", i)
}

func Float64(f float64, precision ...int) ValueFeed {
	return func(v *Value) {
		if v.isSet {
			return
		}
		prec := 0
		for _, p := range precision {
			prec = p
		}
		v.Type = "float64"
		v.Code = formatFloat(f, prec)
		v.ValueFloat64 = f
		v.isSet = true
	}
}

func formatFloat(f float64, precicion int) string {
	return fmt.Sprintf("%v", f)
}

func Description(desc string) ValueFeed {
	return func(v *Value) {
		v.ValueDescription = desc
	}
}

func Separator(omitable bool) ValueFeed {
	return func(v *Value) {
		v.Code = "-"
		v.Key = "-"
		v.Type = "separator"
		v.isSet = true
		v.Omitable = omitable
	}
}
