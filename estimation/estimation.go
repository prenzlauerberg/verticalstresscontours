package estimation

import (
	"CE366VerticalStress/functions"
	"CE366VerticalStress/types"
	"fmt"
)

type Estimator struct {
	Q            float64
	B            float64
	CurrentDepth float64
	Curves       []*types.CurveInf
	sym          bool
	Trace        []types.Direction
}

func (estimator *Estimator) CalculateTheContours() {
	for curveIndex, curve := range estimator.Curves {
		goal := estimator.Q * curve.Coefficient
		y := (estimator.B / 2.0) - 0.02*estimator.B
		z := 0.01
		var iteration = true
		var iterationIndex = 0
		for iteration {
			iterationIndex++
			//fmt.Printf("iteration index %d , y is %v z is %v ,\n", iterationIndex, y, z)
			estimated, _, found := estimator.EstimateSigmaV(goal, &y, &z, curveIndex)
			if estimated {
				fmt.Printf("estimation concluded the goal is, %v , found %v , stopping iteration\n", goal, found)
				z = z + 0.008*estimator.B
			}
			if estimator.sym {
				iteration = false
			}

		}

	}
}

func (estimator *Estimator) EstimateSigmaV(goal float64, y *float64, z *float64, curveIndex int) (bool, float64, float64) {
	sigmaV := functions.CalculateSigmaVFromCoordinates(estimator.Q, estimator.B, *y, *z)
	result, direction := functions.CheckDeviation(goal, sigmaV)
	if result {
		//fmt.Printf("check deviation returned true\n")
		if functions.CheckIfItsOnBottom(estimator.B, *y, *z) {
			//fmt.Printf("found the bottom\n")
			estimator.sym = true
		}
		estimator.AddPointToTheSlice(curveIndex, *y, *z, sigmaV)
		return true, goal, sigmaV
	}
	functions.NextY(y, direction, estimator.Trace)
	//fmt.Printf("#2->final y is %v\n\n", *y)
	if estimator.Trace != nil {
		estimator.Trace = append(estimator.Trace, direction)
	} else {
		traces := make([]types.Direction, 0)
		traces = append(traces, direction)
		estimator.Trace = traces
	}

	if *y < 0.002 {
		//fmt.Printf("#3->enough y is %v\n\n", *y)

		estimator.sym = true
		estimator.AddPointToTheSlice(curveIndex, *y, *z, sigmaV)
		return true, goal, sigmaV
	}
	//this is fallback
	return false, goal, sigmaV
}

func (estimator *Estimator) AddPointToTheSlice(index int, y float64, z float64, value float64) {
	r := types.RelativePointInfo{
		Y:     y,
		Z:     z,
		Value: value,
	}
	c := estimator.Curves[index]
	c.Points = append(c.Points, r)
}
