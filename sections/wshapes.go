package sections

import (
	"encoding/csv"
	"os"
	"strconv"
)

type WShape struct {
	Shape string
	Zx    float64
	Sx    float64
	Iy    float64
	H0    float64
	J     float64
	Ry    float64
	Cw    float64
	Bf    float64
	Tf    float64
	H     float64
	Tw    float64
}

func parseFloat(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
}

func LoadShapeOptions() ([]string, error) {
	file, err := os.Open("sections/wshapes.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var shapes []string
	for i, record := range records {
		if i == 0 {
			continue //Skip header
		}
		// Shape name in first column
		shapes = append(shapes, record[0])
	}

	return shapes, nil
}

func LoadWShapes() ([]WShape, error) {
	file, err := os.Open("sections/wshapes.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var shapes []WShape
	for i, row := range rows {
		if i == 0 { // skip header
			continue
		}

		s := WShape{
			Shape: row[0],
			Zx:    parseFloat(row[1]),
			Sx:    parseFloat(row[2]),
			Iy:    parseFloat(row[3]),
			H0:    parseFloat(row[4]),
			J:     parseFloat(row[5]),
			Ry:    parseFloat(row[6]),
			Cw:    parseFloat(row[7]),
			Bf:    parseFloat(row[8]),
			Tf:    parseFloat(row[9]),
			H:     parseFloat(row[10]),
			Tw:    parseFloat(row[11]),
		}

		shapes = append(shapes, s)
	}

	return shapes, nil
}
