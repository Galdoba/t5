package profile

import (
	"fmt"
	"strings"

	"github.com/Galdoba/t5/pkg/profile/prfvalue"
)

type UniversalProfile struct {
	Keys   map[int]string
	Values map[int]*prfvalue.Value
	//KeySequance []string
	Err error
}

type Profiler interface {
	Schema() (string, []*prfvalue.Value)
	InjectValue(val *prfvalue.Value) error
}

func Profile(p Profiler) *UniversalProfile {
	usp := UniversalProfile{}
	usp.Keys = make(map[int]string)
	usp.Values = make(map[int]*prfvalue.Value)
	schema, values := p.Schema()
	for i, val := range values {
		k := val.Key
		if !strings.HasPrefix(schema, k) {
			usp.Err = fmt.Errorf("not found value %v (key=%v)", i, k)
			fmt.Println(usp.Err)
			return &usp
		}
		usp.Keys[i] = k
		usp.Values[i] = val
		schema = strings.TrimPrefix(schema, k)
	}
	return &usp
}

func (up *UniversalProfile) String() string {
	max := len(up.Values)
	s := ""
	for i := 0; i < max; i++ {
		s += up.Values[i].Code
	}
	return s
}

func InjectString(p Profiler, feed string) error {
	_, values := p.Schema()
	//codeLenBefore := 0
	injections := []*prfvalue.Value{}
	for _, val := range values {
		valueLen := len(strings.Split(val.Code, ""))
		if valueLen > len(feed) {
			return fmt.Errorf("injection feed out of bounds")
		}
		injection := prfvalue.New(val.Key)
		injection.Type = val.Type
		injection.Code = strings.Join(strings.Split(feed, "")[0:valueLen], "")
		injections = prfvalue.AddValue(injections, injection)
		feed = strings.Join(strings.Split(feed, "")[valueLen:], "")
	}
	if len(strings.Split(feed, "")) != 0 {
		return fmt.Errorf("injection feed not used up")
	}
	for _, injection := range injections {
		if err := p.InjectValue(injection); err != nil {
			return fmt.Errorf("failed injection(%v): %v", injection, err)
		}
	}
	return nil
}

/*
UWP:= StSAHPGL-T
UPP:= SDEIES

*/
