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
}

func parseFloat(s string) float64 {
	v, _ := strconv.ParseFloat(s, 64)
	return v
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
		}

		shapes = append(shapes, s)
	}

	return shapes, nil
}
