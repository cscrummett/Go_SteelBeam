package main

import (
	"github.com/cscrummett/Go_SteelBeam/design"
	"github.com/cscrummett/Go_SteelBeam/sections"
)

func main() {

	var shapeName string = "W10X33"
	var Cb float64 = 1.0     //conservative
	var Fy float64 = 50      //ksi
	var E float64 = 29000    //ksi
	var Lb float64 = 10 * 12 //in unbraced length

	shapes, err := sections.LoadWShapes()
	if err != nil {
		panic(err)
	}

	var shape sections.WShape
	found := false
	for _, s := range shapes {
		if s.Shape == shapeName {
			shape = s
			found = true
			break
		}
	}

	if !found {
		panic("Shape not found: " + shapeName)
	}

	var Mp float64 = design.Mn_Calc(shape, Cb, Fy, E, Lb)
	var Mp_kft float64 = Mp / 12

	print(shapeName, " Final Capacity = ", 0.9*Mp_kft)

}
