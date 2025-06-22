package coordinates

import "testing"

func Test_localToGlobal(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		sx    int
		sy    int
		lx    int
		ly    int
		want  int
		want2 int
	}{
		{name: "Reference", sx: 0, sy: 0, lx: 1, ly: 40, want: 0, want2: 0},
		{name: "Capital", sx: 0, sy: 0, lx: 21, ly: 18, want: 20, want2: -22},
		{name: "Regina", sx: -4, sy: -1, lx: 19, ly: 10, want: -110, want2: -70},
		{name: "Tepav", sx: 4, sy: 2, lx: 9, ly: 8, want: 136, want2: 48},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := localToGlobal(tt.sx, tt.sy, tt.lx, tt.ly)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("localToGlobal(X) = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("localToGlobal(Y) = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_localToCube(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		sx    int
		sy    int
		lx    int
		ly    int
		want  int
		want2 int
		want3 int
	}{
		{name: "Reference", sx: 0, sy: 0, lx: 1, ly: 40, want: 0, want2: 0, want3: 0},
		{name: "Wonderland", sx: -1, sy: 1, lx: 30, ly: 1, want: -3, want2: 3, want3: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2, got3 := localToCube(tt.sx, tt.sy, tt.lx, tt.ly)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("localToCube(Q) = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("localToCube(R) = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("localToCube(S) = %v, want %v", got3, tt.want3)
			}
		})
	}
}
