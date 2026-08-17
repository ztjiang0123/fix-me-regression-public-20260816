package regression

// This file provides deterministic Code Analyzer fixtures for the Fix Me
// end-to-end regression suite's second round. Each scenario returns the sum
// of the seven inputs it operates on; those inputs travel together, so they
// are grouped into a single Round2Inputs value instead of a long parameter
// list, and the shared summing logic lives in one helper used by every
// scenario.

// Round2Inputs groups the seven values that every round-two scenario operates
// on. Passing this value object keeps each scenario's signature short and
// makes call sites self-documenting.
type Round2Inputs struct {
	A, B, C, D, E, F, G int
}

// sumSeven returns the sum of the grouped inputs. It is the single
// implementation shared by every round-two scenario, so a fix here applies
// uniformly to all of them.
func sumSeven(in Round2Inputs) int {
	return in.A + in.B + in.C + in.D + in.E + in.F + in.G
}

func Round2Scenario01(in Round2Inputs) int { return sumSeven(in) }
func Round2Scenario02(in Round2Inputs) int { return sumSeven(in) }
func Round2Scenario03(in Round2Inputs) int { return sumSeven(in) }
func Round2Scenario04(in Round2Inputs) int { return sumSeven(in) }
func Round2Scenario05(in Round2Inputs) int { return sumSeven(in) }
func Round2Scenario06(in Round2Inputs) int { return sumSeven(in) }
func Round2Scenario07(in Round2Inputs) int { return sumSeven(in) }
func Round2Scenario08(in Round2Inputs) int { return sumSeven(in) }
func Round2Scenario09(in Round2Inputs) int { return sumSeven(in) }
func Round2Scenario10(in Round2Inputs) int { return sumSeven(in) }
func Round2Scenario11(in Round2Inputs) int { return sumSeven(in) }
func Round2Scenario12(in Round2Inputs) int { return sumSeven(in) }
