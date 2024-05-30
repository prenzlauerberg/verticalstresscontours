package main

import (
	"CE366VerticalStress/estimation"
	"CE366VerticalStress/types"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
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

	app := fiber.New(fiber.Config{
		AppName: "CE366 Vertical Stress Contours",
	})

	app.Static("/ce366", "./public")

	app.Post("/api/contours", func(c *fiber.Ctx) error {
		requestParameters := new(types.RequestParameters)
		LineMap = make(map[string]*types.Line)
		if err := c.BodyParser(requestParameters); err != nil {
			return err
		}
		qP, err := strconv.ParseFloat(requestParameters.Q, 64)
		if err != nil {
			return err
		}
		bP, err := strconv.ParseFloat(requestParameters.B, 64)
		if err != nil {
			return err
		}
		parameters := types.Parameters{
			Q: qP,
			B: bP,
		}
		log.Info("will calculate with the following parameters", parameters)
		jsonBytes, err := ProcessContours(parameters)
		if err != nil {
			return err
		}
		return c.JSON(string(jsonBytes))
	})

	app.Listen(":9924")

}

func ProcessContours(p types.Parameters) ([]byte, error) {
	//if err := f.SaveAs("points.xlsx"); err != nil { }
	var contours []types.Contour

	// Populate the slice with Contour elements
	for i := 2; i <= 9; i++ {
		coefficient := float64(i) / 10
		name := "Contour-" + strconv.FormatFloat(coefficient, 'f', 1, 64)
		var color string
		if i%9 == 0 {
			color = "black"
		} else if absInt(i-9) == 1 {
			color = "indigo"
		} else if absInt(i-9) == 2 {
			color = "steelblue"
		} else if absInt(i-9) == 3 {
			color = "slategrey"
		} else if absInt(i-9) == 4 {
			color = "forestgreen"
		} else if absInt(i-9) == 5 {
			color = "red"
		} else if absInt(i-9) == 6 {
			color = "green"
		} else if absInt(i-9) == 7 {
			color = "midnightblue"
		} else if absInt(i-9) == 8 {
			color = "red"
		}
		contours = append(contours, types.Contour{Name: name, Coefficient: coefficient, Color: color})
	}

	// Iterate over the slice and print each element
	for _, contour := range contours {
		//fmt.Printf("Name: %s, Coefficient: %.1f\n", contour.Name, contour.Coefficient)
		ProcessCurve(contour.Name, contour.Coefficient, p, contour.Color)
	}
	return LineMapToJSON()
}

func ProcessCurve(name string, coeff float64, p types.Parameters, color string) {
	curveList := make([]*types.CurveInf, 0)
	curve00 := types.CurveInf{Name: name, Coefficient: coeff}
	curveList = append(curveList, &curve00)
	est := estimation.Estimator{Q: p.Q, B: p.B, Curves: curveList}
	est.CalculateTheContours()

	for _, curve := range est.Curves {
		for _, dataPoint := range curve.Points {
			SavePointToLine(curve.Name, dataPoint, color)
		}
	}
}

func SavePointToLine(name string, point types.RelativePointInfo, color string) {
	n, ok := LineMap[name]
	if ok {
		point1 := types.Point{
			X: point.Y,
			Y: point.Z * -1.0,
		}
		point2 := types.Point{
			X: point.Y * -1.0,
			Y: point.Z * -1.0,
		}
		n.Points = append(n.Points, point1, point2)

	} else {
		LineMap[name] = &types.Line{
			Color:  color,
			Name:   name,
			Points: []types.Point{},
		}
		s := LineMap[name]
		point0 := types.Point{
			X: point.Y,
			Y: point.Z * -1.0,
		}
		point02 := types.Point{
			X: point.Y * -1.0,
			Y: point.Z * -1.0,
		}
		s.Points = append(s.Points, point0, point02)
	}
}

func LineMapToJSON() ([]byte, error) {
	// Convert lines to JSON
	linesJSON, err := json.Marshal(LineMap)
	if err != nil {
		fmt.Println("Error marshaling lines:", err)
		return nil, errors.New(err.Error())
	}
	return linesJSON, nil
	/*// Create and write to points.js file
	file, err := os.Create("points.js")
	if err != nil {
		fmt.Println("Error creating points.js file:", err)
		return nil, errors.New( err.Error())
	}
	defer file.Close()

	jsContent := fmt.Sprintf("window.lines = %s;", linesJSON)
	_, err = file.WriteString(jsContent)
	if err != nil {
		fmt.Println("Error writing to points.js file:", err)
		return nil, errors.New( err.Error())
	}
	*/
}

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
