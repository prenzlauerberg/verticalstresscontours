package main

import (
	"CE366VerticalStress/estimation"
	"CE366VerticalStress/types"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var LineMap map[string]*types.Line

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
	LineMap = make(map[string]*types.Line)
}

func main() {
	//f := excelize.NewFile()
	//f.SetCellValue("Sheet1", "A1", "Y")
	//f.SetCellValue("Sheet1", "B1", "Z")
	//f.SetCellValue("Sheet1", "C1", "Value")

	//if err := f.SaveAs("points.xlsx"); err != nil { }
	var contours []types.Contour

	// Populate the slice with Contour elements
	for i := 1; i <= 9; i++ {
		coefficient := float64(i) / 10
		name := "Contour-" + strconv.FormatFloat(coefficient, 'f', 1, 64)
		contours = append(contours, types.Contour{Name: name, Coefficient: coefficient})
	}

	// Iterate over the slice and print each element
	for _, contour := range contours {
		//fmt.Printf("Name: %s, Coefficient: %.1f\n", contour.Name, contour.Coefficient)
		ProcessCurve(contour.Name, contour.Coefficient)
	}
	SaveLinesToJS()

}

func ProcessCurve(name string, coeff float64) {
	curveList := make([]*types.CurveInf, 0)
	curve00 := types.CurveInf{Name: name, Coefficient: coeff}
	curveList = append(curveList, &curve00)
	est := estimation.Estimator{Q: 12, B: 10, Curves: curveList}
	est.CalculateTheContours()

	for _, curve := range est.Curves {
		for _, dataPoint := range curve.Points {
			SavePointToLine(curve.Name, dataPoint)
		}
	}
}

func SavePointToLine(name string, point types.RelativePointInfo) {
	n, ok := LineMap[name]
	if ok {
		point1 := types.Point{
			X: point.Y,
			Y: point.Z,
		}
		point2 := types.Point{
			X: point.Y * -1.0,
			Y: point.Z,
		}
		n.Points = append(n.Points, point1, point2)

	} else {
		LineMap[name] = &types.Line{
			Color:  "",
			Name:   name,
			Points: []types.Point{},
		}
		s := LineMap[name]
		point0 := types.Point{
			X: point.Y,
			Y: point.Z,
		}
		point02 := types.Point{
			X: point.Y * -1.0,
			Y: point.Z,
		}
		s.Points = append(s.Points, point0, point02)
	}
}

func SaveLinesToJS() {
	// Convert lines to JSON
	linesJSON, err := json.Marshal(LineMap)
	if err != nil {
		fmt.Println("Error marshaling lines:", err)
		return
	}
	// Create and write to points.js file
	file, err := os.Create("points.js")
	if err != nil {
		fmt.Println("Error creating points.js file:", err)
		return
	}
	defer file.Close()

	jsContent := fmt.Sprintf("window.lines = %s;", linesJSON)
	_, err = file.WriteString(jsContent)
	if err != nil {
		fmt.Println("Error writing to points.js file:", err)
	}

	fmt.Println("points.js file generated successfully")
}
