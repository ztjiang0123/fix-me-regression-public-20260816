package regression

import "testing"

// sample exercises every scenario with distinct values so the pair each
// scenario selects is unambiguous.
var sample = Inputs{
	A: 1, B: 2, C: 3, D: 4, E: 5, F: 6,
	G: 7, H: 8, I: 9, J: 10, K: 11, L: 12,
}

func TestScenarios(t *testing.T) {
	cases := []struct {
		name string
		got  int
		want int
	}{
		{"Scenario01", Scenario01(sample), 1 + 12},
		{"Scenario02", Scenario02(sample), 2 + 11},
		{"Scenario03", Scenario03(sample), 3 + 10},
		{"Scenario04", Scenario04(sample), 4 + 9},
		{"Scenario05", Scenario05(sample), 5 + 8},
		{"Scenario06", Scenario06(sample), 6 + 7},
		{"Scenario07", Scenario07(sample), 7 + 6},
		{"Scenario08", Scenario08(sample), 8 + 5},
		{"Scenario09", Scenario09(sample), 9 + 4},
		{"Scenario10", Scenario10(sample), 10 + 3},
		{"Scenario11", Scenario11(sample), 11 + 2},
		{"Scenario12", Scenario12(sample), 12 + 1},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("%s = %d, want %d", c.name, c.got, c.want)
		}
	}
}

func TestCurrentRevision(t *testing.T) {
	if got := CurrentRevision(); got != 2 {
		t.Errorf("CurrentRevision() = %d, want 2", got)
	}
}
