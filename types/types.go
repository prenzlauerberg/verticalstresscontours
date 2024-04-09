package types

type Direction int

var (
	UP               Direction = 0
	DOWN             Direction = 1
	STOP             Direction = -1
	RIGHT            Direction = 90
	LEFT             Direction = -90
	AllowedDeviation float64   = 0.001
)

type RelativePointInfo struct {
	Y     float64
	Z     float64
	Value float64
}

type CurveInf struct {
	Name        string
	Coefficient float64
	Points      []RelativePointInfo
}
