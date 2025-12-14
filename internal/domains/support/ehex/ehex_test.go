package ehex

import (
	"testing"
)

// TestFromValueValid tests FromValue with all valid integer inputs (0-35).
func TestFromValueValid(t *testing.T) {
	testCases := []struct {
		name     string
		input    int
		expected string
	}{
		{"Zero", 0, "0"},
		{"One", 1, "1"},
		{"Nine", 9, "9"},
		{"Ten", 10, "A"},
		{"Eleven", 11, "B"},
		{"Fifteen", 15, "F"},
		{"Sixteen", 16, "G"},
		{"Seventeen", 17, "H"},
		{"Eighteen", 18, "J"},
		{"TwentyTwo", 22, "N"},
		{"TwentyThree", 23, "P"},
		{"ThirtyThree", 33, "Z"},
		{"Unknown", 34, "?"},
		{"Any", 35, "*"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			eh := FromValue(tc.input)
			if eh.Code() != tc.expected {
				t.Errorf("FromValue(%d) = %s, want %s", tc.input, eh.Code(), tc.expected)
			}
			if eh.Value() != tc.input {
				t.Errorf("FromValue(%d).Value() = %d, want %d", tc.input, eh.Value(), tc.input)
			}
		})
	}
}

// TestFromCodeValid tests FromCode with all valid string inputs.
func TestFromCodeValid(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"Zero", "0", 0},
		{"One", "1", 1},
		{"Nine", "9", 9},
		{"A", "A", 10},
		{"B", "B", 11},
		{"F", "F", 15},
		{"G", "G", 16},
		{"H", "H", 17},
		{"J", "J", 18},
		{"N", "N", 22},
		{"P", "P", 23},
		{"Z", "Z", 33},
		{"Unknown", "?", 34},
		{"Any", "*", 35},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			eh := FromCode(tc.input)
			if eh.Code() != tc.input {
				t.Errorf("FromCode(%s) = %s, want %s", tc.input, eh.Code(), tc.input)
			}
			if eh.Value() != tc.expected {
				t.Errorf("FromCode(%s).Value() = %d, want %d", tc.input, eh.Value(), tc.expected)
			}
		})
	}
}

// TestFromValueInvalid tests that FromValue panics with invalid inputs.
func TestFromValueInvalid(t *testing.T) {
	testCases := []struct {
		name  string
		input int
	}{
		{"Negative", -1},
		{"TooLarge", 36},
		{"VeryLarge", 100},
		{"VeryNegative", -100},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("FromValue(%d) did not panic", tc.input)
				}
			}()
			FromValue(tc.input)
		})
	}
}

// TestFromCodeInvalid tests that FromCode panics with invalid inputs.
func TestFromCodeInvalid(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{"Empty", ""},
		{"InvalidChar", "I"},
		{"InvalidChar", "O"},
		{"LowerCase", "a"},
		{"Symbol", "$"},
		{"MultiChar", "AA"},
		{"Space", " "},
		{"Newline", "\n"},
		{"Tab", "\t"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("FromCode(%q) did not panic", tc.input)
				}
			}()
			FromCode(tc.input)
		})
	}
}

// TestCodeMethod tests the Code method for all pre-defined Ehex constants.
func TestCodeMethod(t *testing.T) {
	testCases := []struct {
		name   string
		ehex   *Ehex
		expect string
	}{
		{"0", &Ehex_0, "0"},
		{"1", &Ehex_1, "1"},
		{"A", &Ehex_A, "A"},
		{"Z", &Ehex_Z, "Z"},
		{"Unknown", &Ehex_Unknown, "?"},
		{"Any", &Ehex_Any, "*"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.ehex.Code(); got != tc.expect {
				t.Errorf("%s.Code() = %s, want %s", tc.name, got, tc.expect)
			}
		})
	}
}

// TestValueMethod tests the Value method for all pre-defined Ehex constants.
func TestValueMethod(t *testing.T) {
	testCases := []struct {
		name   string
		ehex   *Ehex
		expect int
	}{
		{"0", &Ehex_0, 0},
		{"1", &Ehex_1, 1},
		{"A", &Ehex_A, 10},
		{"Z", &Ehex_Z, 33},
		{"Unknown", &Ehex_Unknown, 34},
		{"Any", &Ehex_Any, 35},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.ehex.Value(); got != tc.expect {
				t.Errorf("%s.Value() = %d, want %d", tc.name, got, tc.expect)
			}
		})
	}
}

// TestAllValues tests round-trip conversion for all valid values.
func TestAllValues(t *testing.T) {
	for i := 0; i <= 35; i++ {
		t.Run(string(FromValue(i).Code()), func(t *testing.T) {
			// Test FromValue -> Code -> FromCode -> Value round trip
			eh1 := FromValue(i)
			code := eh1.Code()
			eh2 := FromCode(code)
			value := eh2.Value()

			if value != i {
				t.Errorf("Round trip failed: %d -> %s -> %d", i, code, value)
			}

			// Test that the code is the same after round trip
			if eh2.Code() != code {
				t.Errorf("Code changed after round trip: %s -> %s", code, eh2.Code())
			}
		})
	}
}

// TestPointerEquality tests that FromValue and FromCode return pointers to the same instances.
func TestPointerEquality(t *testing.T) {
	// Test that FromValue returns the same pointer for the same input
	eh1 := FromValue(10)
	eh2 := FromValue(10)
	if eh1 != eh2 {
		t.Errorf("FromValue(10) returned different pointers")
	}

	// Test that FromCode returns the same pointer for the same input
	eh3 := FromCode("A")
	eh4 := FromCode("A")
	if eh3 != eh4 {
		t.Errorf("FromCode('A') returned different pointers")
	}

	// Test that FromValue and FromCode return the same pointer for equivalent values
	if eh1 != eh3 {
		t.Errorf("FromValue(10) and FromCode('A') returned different pointers")
	}
}

// TestNilChecks tests internal helper functions with nil returns.
func TestNilChecks(t *testing.T) {
	// Test newEhexInt with invalid value
	if eh := newEhexInt(100); eh != nil {
		t.Errorf("newEhexInt(100) should return nil, got %v", eh)
	}

	// Test newEhexString with invalid code
	if eh := newEhexString("I"); eh != nil {
		t.Errorf("newEhexString('I') should return nil, got %v", eh)
	}
}
