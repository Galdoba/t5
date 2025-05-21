package profile

// const (
// 	Field_Code        = "code"
// 	Field_Type        = "type"
// 	Field_Description = "description"
// )

// type ProfileValue struct {
// 	Type             string  `json:"type,omitempty"`
// 	Code             string  `json:"Code,omitempty"`
// 	ValueInt         int     `json:"ValueInt,omitempty"`
// 	ValueFloat64     float64 `json:"ValueFloat64,omitempty"`
// 	ValueBool        bool    `json:"ValueBool,omitempty"`
// 	ValueDescription string  `json:"ValueDescription,omitempty"`
// }

// type ValueProvider interface {
// 	Provide(string) string
// }

// type Feed interface {
// 	int | float64 | bool
// }

// func NewValue(feed ValueProvider) *ProfileValue {
// 	pv := ProfileValue{}
// 	pv.Code = feed.Provide(Field_Code)
// 	pv.Type = feed.Provide(Field_Type)
// 	switch pv.Type {
// 	case "separator", "splitter":
// 	case "bool":
// 		f := feed.(ValueBool)
// 		pv.ValueBool = f.Value()
// 	case "int":
// 		f := feed.(ValueInt)
// 		pv.ValueInt = f.Value()
// 	case "float64":
// 		f := feed.(ValueFloat64)
// 		pv.ValueFloat64 = f.Value()
// 	default:
// 		return nil
// 	}
// 	return &pv
// }

// type ValueInt interface {
// 	ValueProvider
// 	Value() int
// }

// type ValueFloat64 interface {
// 	ValueProvider
// 	Value() float64
// }

// type ValueBool interface {
// 	ValueProvider
// 	Value() bool
// }

// /*
// up.UWP()
// */
