package main

import (
	"CE366VerticalStress/estimation"
	"CE366VerticalStress/types"
	log "github.com/sirupsen/logrus"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
)

type file excelize.File

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "Y")
	f.SetCellValue("Sheet1", "B1", "Z")
	f.SetCellValue("Sheet1", "C1", "Value")
	curveList := make([]*types.CurveInf, 0)
	//curve01 := types.CurveInf{Name: "0.1", Coefficient: 0.1}
	curve01 := types.CurveInf{Name: "0.2", Coefficient: 0.2}
	curveList = append(curveList, &curve01)
	est := estimation.Estimator{Q: 12, B: 10, Curves: curveList}
	est.CalculateTheContours()

	for _, curve := range est.Curves {
		for i, dataPoint := range curve.Points {
			SaveRow(f, i, dataPoint)
		}
	}

	if err := f.SaveAs("points.xlsx"); err != nil {

	}

}

func SaveRow(excelFile *excelize.File, i int, point types.RelativePointInfo) {
	excelFile.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), point.Y)
	excelFile.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), point.Z)
	excelFile.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), point.Value)
}
