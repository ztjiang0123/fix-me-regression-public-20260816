package regression

import "testing"

// p is a Params value with distinct field values so that any accidental
// swap between operands is detectable in the assertions below.
var p = Params{A: 1, B: 2, C: 3, D: 4, E: 5, F: 6, G: 7, H: 8, I: 9, J: 10, K: 11, L: 12}

func TestScenarios(t *testing.T) {
	cases := []struct {
		name string
		got  int
		want int
	}{
		{"Scenario01", Scenario01(p), p.A + p.L},
		{"Scenario02", Scenario02(p), p.B + p.K},
		{"Scenario03", Scenario03(p), p.C + p.J},
		{"Scenario04", Scenario04(p), p.D + p.I},
		{"Scenario05", Scenario05(p), p.E + p.H},
		{"Scenario06", Scenario06(p), p.F + p.G},
		{"Scenario07", Scenario07(p), p.G + p.F},
		{"Scenario08", Scenario08(p), p.H + p.E},
		{"Scenario09", Scenario09(p), p.I + p.D},
		{"Scenario10", Scenario10(p), p.J + p.C},
		{"Scenario11", Scenario11(p), p.K + p.B},
		{"Scenario12", Scenario12(p), p.L + p.A},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("%s = %d, want %d", c.name, c.got, c.want)
		}
	}
}
