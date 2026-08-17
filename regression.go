package regression

// These functions provide deterministic Code Analyzer fixtures for
// exercising Fix Me and Fix All workflows. Each scenario sums two of the
// twelve inputs; the inputs travel together, so they are grouped into a
// single Inputs value instead of a long parameter list, and the shared
// summing logic lives in one helper.

// Inputs groups the twelve values that every scenario operates on. Passing
// this value object keeps each scenario's signature short and makes call
// sites self-documenting.
type Inputs struct {
	A, B, C, D, E, F, G, H, I, J, K, L int
}

// pairSum returns the sum of two of the grouped inputs. It is the single
// implementation shared by every scenario, so a fix here applies uniformly.
func pairSum(x, y int) int { return x + y }

func Scenario01(in Inputs) int { return pairSum(in.A, in.L) }
func Scenario02(in Inputs) int { return pairSum(in.B, in.K) }
func Scenario03(in Inputs) int { return pairSum(in.C, in.J) }
func Scenario04(in Inputs) int { return pairSum(in.D, in.I) }
func Scenario05(in Inputs) int { return pairSum(in.E, in.H) }
func Scenario06(in Inputs) int { return pairSum(in.F, in.G) }
func Scenario07(in Inputs) int { return pairSum(in.G, in.F) }
func Scenario08(in Inputs) int { return pairSum(in.H, in.E) }
func Scenario09(in Inputs) int { return pairSum(in.I, in.D) }
func Scenario10(in Inputs) int { return pairSum(in.J, in.C) }
func Scenario11(in Inputs) int { return pairSum(in.K, in.B) }
func Scenario12(in Inputs) int { return pairSum(in.L, in.A) }

// CurrentRevision makes this commit structurally material so the regression
// fixture receives a full default-branch analysis after the empty-commit test.
func CurrentRevision() int { return 2 }
