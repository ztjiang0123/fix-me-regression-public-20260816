package regression

import "testing"

// round2Sample exercises every round-two scenario with distinct values so the
// full sum each scenario computes is unambiguous.
var round2Sample = Round2Inputs{
	A: 1, B: 2, C: 3, D: 4, E: 5, F: 6, G: 7,
}

// wantRound2Sum is the sum of every field in round2Sample. All round-two
// scenarios share the sumInputs helper, so they must all return this value.
const wantRound2Sum = 1 + 2 + 3 + 4 + 5 + 6 + 7

// TestRound2Scenarios verifies the struct-based signature (Round2Inputs)
// preserves the original summing behavior for every scenario. This locks in
// the parameter_count refactor that grouped the seven int parameters into a
// single value object.
func TestRound2Scenarios(t *testing.T) {
	cases := []struct {
		name string
		got  int
	}{
		{"Round2Scenario01", Round2Scenario01(round2Sample)},
		{"Round2Scenario02", Round2Scenario02(round2Sample)},
		{"Round2Scenario03", Round2Scenario03(round2Sample)},
		{"Round2Scenario04", Round2Scenario04(round2Sample)},
		{"Round2Scenario05", Round2Scenario05(round2Sample)},
		{"Round2Scenario06", Round2Scenario06(round2Sample)},
		{"Round2Scenario07", Round2Scenario07(round2Sample)},
		{"Round2Scenario08", Round2Scenario08(round2Sample)},
		{"Round2Scenario09", Round2Scenario09(round2Sample)},
		{"Round2Scenario10", Round2Scenario10(round2Sample)},
		{"Round2Scenario11", Round2Scenario11(round2Sample)},
		{"Round2Scenario12", Round2Scenario12(round2Sample)},
	}
	for _, c := range cases {
		if c.got != wantRound2Sum {
			t.Errorf("%s = %d, want %d", c.name, c.got, wantRound2Sum)
		}
	}
}
