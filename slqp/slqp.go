package slqp

type Status int

const (
	NotTerminated Status = iota
	Success
	FunctionThreshold
	FunctionConvergence
	GradientThreshold
	StepConvergence
	FunctionNegativeInfinity
	MethodConverge
	Failure
	IterationLimit
	RuntimeLimit
	FunctionEvaluationLimit
	GradientEvaluationLimit
	HessianEvaluationLimit
)

type Problem struct {
	Obj              func([]float64) float64
	EqCons, IneqCons func([]float64) []float64
	N int
}

func dot(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("length mismatch")
	}

	var sum float64
	for i := range a {
		sum += a[i] * b[i]
	}

	return sum
}

func (p Problem) Lagrangian(x, lmd, sig []float64) float64 {
	return p.Obj(x) - dot(lmd, p.IneqCons(x)) - dot(sig, p.EqCons(x))
}

func (p Problem) Minimize() ([]float64, float64) {
	// TODO: Implement
	var optX []float64
	var cost float64
	return optX, cost
}
