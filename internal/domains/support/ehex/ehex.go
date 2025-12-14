// Package ehex provides Extended Hexadecimal (Ehex) encoding and decoding functionality.
// Ehex is a base-36 encoding system used in Traveller and other tabletop RPGs,
// which extends hexadecimal with additional characters to represent values 0-35.
// The encoding excludes 'I' and 'O' to avoid confusion with numbers 1 and 0.
//
// The package defines constants for all valid Ehex codes and provides
// conversion functions between integer values, string codes, and Ehex objects.
package ehex

import "fmt"

// ehexCode represents the string representation of an Ehex value.
// Valid codes include: 0-9, A-H, J-N, P-Z, ?, and *.
// '?' represents an unknown value, '*' represents "any" value.
type ehexCode string

// ehexValue represents the integer value of an Ehex code.
// Valid values range from 0 to 35 inclusive.
// Special values: 34 = unknown (?), 35 = any (*).
type ehexValue int

// Ehex represents a complete Extended Hexadecimal entity.
// It encapsulates both the string code and integer value representations.
// The zero value is not valid; use the package functions to create Ehex instances.
type Ehex struct {
	// code holds the string representation of the Ehex value
	code ehexCode
	// value holds the integer representation of the Ehex value
	value ehexValue
}

// Package constants for all valid Ehex values.
// These provide pre-allocated instances for each valid Ehex code/value pair.
var (
	Ehex_0       = Ehex{"0", 0}
	Ehex_1       = Ehex{"1", 1}
	Ehex_2       = Ehex{"2", 2}
	Ehex_3       = Ehex{"3", 3}
	Ehex_4       = Ehex{"4", 4}
	Ehex_5       = Ehex{"5", 5}
	Ehex_6       = Ehex{"6", 6}
	Ehex_7       = Ehex{"7", 7}
	Ehex_8       = Ehex{"8", 8}
	Ehex_9       = Ehex{"9", 9}
	Ehex_A       = Ehex{"A", 10}
	Ehex_B       = Ehex{"B", 11}
	Ehex_C       = Ehex{"C", 12}
	Ehex_D       = Ehex{"D", 13}
	Ehex_E       = Ehex{"E", 14}
	Ehex_F       = Ehex{"F", 15}
	Ehex_G       = Ehex{"G", 16}
	Ehex_H       = Ehex{"H", 17}
	Ehex_J       = Ehex{"J", 18}
	Ehex_K       = Ehex{"K", 19}
	Ehex_L       = Ehex{"L", 20}
	Ehex_M       = Ehex{"M", 21}
	Ehex_N       = Ehex{"N", 22}
	Ehex_P       = Ehex{"P", 23}
	Ehex_Q       = Ehex{"Q", 24}
	Ehex_R       = Ehex{"R", 25}
	Ehex_S       = Ehex{"S", 26}
	Ehex_T       = Ehex{"T", 27}
	Ehex_U       = Ehex{"U", 28}
	Ehex_V       = Ehex{"V", 29}
	Ehex_W       = Ehex{"W", 30}
	Ehex_X       = Ehex{"X", 31}
	Ehex_Y       = Ehex{"Y", 32}
	Ehex_Z       = Ehex{"Z", 33}
	Ehex_Unknown = Ehex{"?", 34} // Represents an unknown or undefined value
	Ehex_Any     = Ehex{"*", 35} // Represents a wildcard or "any" value
)

// newEhexInt creates an Ehex pointer from an integer value.
// Returns a pointer to the corresponding pre-allocated Ehex instance,
// or nil if the integer is outside the valid range (0-35).
// This is an internal helper function; use FromValue for public API.
func newEhexInt(i int) *Ehex {
	switch i {
	case 0:
		return &Ehex_0
	case 1:
		return &Ehex_1
	case 2:
		return &Ehex_2
	case 3:
		return &Ehex_3
	case 4:
		return &Ehex_4
	case 5:
		return &Ehex_5
	case 6:
		return &Ehex_6
	case 7:
		return &Ehex_7
	case 8:
		return &Ehex_8
	case 9:
		return &Ehex_9
	case 10:
		return &Ehex_A
	case 11:
		return &Ehex_B
	case 12:
		return &Ehex_C
	case 13:
		return &Ehex_D
	case 14:
		return &Ehex_E
	case 15:
		return &Ehex_F
	case 16:
		return &Ehex_G
	case 17:
		return &Ehex_H
	case 18:
		return &Ehex_J
	case 19:
		return &Ehex_K
	case 20:
		return &Ehex_L
	case 21:
		return &Ehex_M
	case 22:
		return &Ehex_N
	case 23:
		return &Ehex_P
	case 24:
		return &Ehex_Q
	case 25:
		return &Ehex_R
	case 26:
		return &Ehex_S
	case 27:
		return &Ehex_T
	case 28:
		return &Ehex_U
	case 29:
		return &Ehex_V
	case 30:
		return &Ehex_W
	case 31:
		return &Ehex_X
	case 32:
		return &Ehex_Y
	case 33:
		return &Ehex_Z
	case 34:
		return &Ehex_Unknown
	case 35:
		return &Ehex_Any
	default:
		return nil
	}
}

// newEhexString creates an Ehex pointer from a string code.
// Returns a pointer to the corresponding pre-allocated Ehex instance,
// or nil if the string is not a valid Ehex code.
// This is an internal helper function; use FromCode for public API.
func newEhexString(s string) *Ehex {
	switch s {
	case "0":
		return &Ehex_0
	case "1":
		return &Ehex_1
	case "2":
		return &Ehex_2
	case "3":
		return &Ehex_3
	case "4":
		return &Ehex_4
	case "5":
		return &Ehex_5
	case "6":
		return &Ehex_6
	case "7":
		return &Ehex_7
	case "8":
		return &Ehex_8
	case "9":
		return &Ehex_9
	case "A":
		return &Ehex_A
	case "B":
		return &Ehex_B
	case "C":
		return &Ehex_C
	case "D":
		return &Ehex_D
	case "E":
		return &Ehex_E
	case "F":
		return &Ehex_F
	case "G":
		return &Ehex_G
	case "H":
		return &Ehex_H
	case "J":
		return &Ehex_J
	case "K":
		return &Ehex_K
	case "L":
		return &Ehex_L
	case "M":
		return &Ehex_M
	case "N":
		return &Ehex_N
	case "P":
		return &Ehex_P
	case "Q":
		return &Ehex_Q
	case "R":
		return &Ehex_R
	case "S":
		return &Ehex_S
	case "T":
		return &Ehex_T
	case "U":
		return &Ehex_U
	case "V":
		return &Ehex_V
	case "W":
		return &Ehex_W
	case "X":
		return &Ehex_X
	case "Y":
		return &Ehex_Y
	case "Z":
		return &Ehex_Z
	case "?":
		return &Ehex_Unknown
	case "*":
		return &Ehex_Any
	default:
		return nil
	}
}

// FromValue creates an Ehex instance from an integer value.
// Valid input range is 0 to 35 inclusive.
// Panics if the input value is outside the valid range.
// Returns a pointer to a pre-allocated Ehex instance.
func FromValue(value int) *Ehex {
	e := newEhexInt(value)
	if e == nil {
		panic(fmt.Sprintf("'%d' is not ehex value", value))
	}
	return e
}

// FromCode creates an Ehex instance from a string code.
// Valid codes: "0"-"9", "A"-"H", "J"-"N", "P"-"Z", "?", "*".
// Panics if the input string is not a valid Ehex code.
// Returns a pointer to a pre-allocated Ehex instance.
func FromCode(code string) *Ehex {
	e := newEhexString(code)
	if e == nil {
		panic(fmt.Sprintf("'%s' is not ehex code", code))
	}
	return e
}

// Code returns the string representation of the Ehex value.
// This is the human-readable form of the Ehex code.
func (eh *Ehex) Code() string {
	return string(eh.code)
}

// Value returns the integer representation of the Ehex value.
// This is the numeric form of the Ehex code.
func (eh *Ehex) Value() int {
	return int(eh.value)
}
