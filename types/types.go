package types

type Direction int

var (
	UP               Direction = 0
	DOWN             Direction = 1
	STOP             Direction = -1
	RIGHT            Direction = 90
	LEFT             Direction = -90
	AllowedDeviation float64   = 0.0002
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

// Point represents a 2D point ( this will be stored in the Line array/slice )
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"` // values will be mapped from "z" to "y"
}

// Line represents a set of points for a single line (the contour line in this case)
type Line struct {
	Color  string  `json:"color"`  // Optional: color for the line
	Name   string  `json:"name"`   //  Name for the line
	Points []Point `json:"points"` // Points in the line
}

// Contour represents a structure with a name and coefficient
type Contour struct {
	Name        string
	Coefficient float64
}
