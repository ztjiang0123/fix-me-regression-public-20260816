package regression

// These functions provide deterministic Code Analyzer fixtures for
// exercising Fix Me and Fix All workflows. Each scenario sums two of the
// twelve operands; the operands travel together, so they are grouped into a
// single Params value instead of a long parameter list, and the shared
// summing logic lives in one helper rather than being copied per scenario.

// Params groups the twelve operands that every scenario operates on.
type Params struct {
	A, B, C, D, E, F, G, H, I, J, K, L int
}

// sum returns x + y. It is the single implementation shared by every
// scenario so the addition logic is defined in exactly one place.
func sum(x, y int) int { return x + y }

func Scenario01(p Params) int { return sum(p.A, p.L) }
func Scenario02(p Params) int { return sum(p.B, p.K) }
func Scenario03(p Params) int { return sum(p.C, p.J) }
func Scenario04(p Params) int { return sum(p.D, p.I) }
func Scenario05(p Params) int { return sum(p.E, p.H) }
func Scenario06(p Params) int { return sum(p.F, p.G) }
func Scenario07(p Params) int { return sum(p.G, p.F) }
func Scenario08(p Params) int { return sum(p.H, p.E) }
func Scenario09(p Params) int { return sum(p.I, p.D) }
func Scenario10(p Params) int { return sum(p.J, p.C) }
func Scenario11(p Params) int { return sum(p.K, p.B) }
func Scenario12(p Params) int { return sum(p.L, p.A) }
