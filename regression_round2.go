package regression

// This file provides deterministic Code Analyzer fixtures for the Fix Me
// end-to-end regression suite. The seven values every scenario operates on
// travel together, so they are grouped into a single Round2Inputs value
// instead of a long parameter list, and the shared summing logic lives in
// one helper.

// Round2Inputs groups the seven values that every round-two scenario sums.
// Passing this value object keeps each scenario's signature short and makes
// call sites self-documenting.
type Round2Inputs struct {
	A, B, C, D, E, F, G int
}

// sumInputs returns the sum of all grouped inputs. It is the single
// implementation shared by every scenario, so a fix here applies uniformly.
func sumInputs(in Round2Inputs) int {
	return in.A + in.B + in.C + in.D + in.E + in.F + in.G
}

func Round2Scenario01(in Round2Inputs) int { return sumInputs(in) }
func Round2Scenario02(in Round2Inputs) int { return sumInputs(in) }
func Round2Scenario03(in Round2Inputs) int { return sumInputs(in) }
func Round2Scenario04(in Round2Inputs) int { return sumInputs(in) }
func Round2Scenario05(in Round2Inputs) int { return sumInputs(in) }
func Round2Scenario06(in Round2Inputs) int { return sumInputs(in) }
func Round2Scenario07(in Round2Inputs) int { return sumInputs(in) }
func Round2Scenario08(in Round2Inputs) int { return sumInputs(in) }
func Round2Scenario09(in Round2Inputs) int { return sumInputs(in) }
func Round2Scenario10(in Round2Inputs) int { return sumInputs(in) }
func Round2Scenario11(in Round2Inputs) int { return sumInputs(in) }
func Round2Scenario12(in Round2Inputs) int { return sumInputs(in) }
