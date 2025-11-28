package design

import (
	"fmt"
	"math"

	"github.com/cscrummett/Go_SteelBeam/sections"
)

// always calc this (Section F2)
func Mn_Calc(shape sections.WShape, Cb float64, Fy float64, E float64, Lb float64) float64 {
	var Lp float64 = 1.76 * shape.Ry * math.Sqrt(E/Fy)         //F2-5 (in) Limiting laterally unbraced length for yielding limit state
	var jcsxh0 float64 = shape.J * 1.0 / (shape.Sx * shape.H0) //c = 1.0 for W-flange (F2-8a)
	var rts float64 = math.Sqrt(math.Sqrt(shape.Iy*shape.Cw) / shape.Sx)

	var Mp float64 = Fy * shape.Zx // k-in
	var Mp_kft float64 = Mp / 12   //k-ft

	fmt.Printf("Lb = %.1f ft\n", Lb/12)
	fmt.Printf("Lp = %.1f ft\n", Lp/12)

	if Lb <= Lp { //case a (plastic hinge)
		fmt.Println("Case A: Limit State is Plastic Hinge")
		fmt.Printf("Mp = %.0f k-ft\n", Mp_kft)
		return Mp //return Mp
	}

	var Lr_01 float64 = math.Sqrt(jcsxh0 + math.Sqrt(math.Pow(jcsxh0, 2)+6.76*math.Pow(0.7*Fy/E, 2)))
	var Lr float64 = 1.95 * rts * E / (0.7 * Fy) * Lr_01

	if Lb > Lr { //case c (elastic LTB)
		var lbrts2 float64 = math.Pow(Lb/rts, 2)
		var Fcr_01 float64 = (Cb * math.Pow(math.Pi, 2) * E) / (lbrts2)
		var Fcr float64 = Fcr_01 * math.Sqrt(1+0.078*jcsxh0*lbrts2) //critical stress, ksi
		var Mn float64 = Fcr * shape.Sx                             //k-in
		var Mn_kft float64 = Mn / 12                                //k-ft
		fmt.Println("Case C: Limit State is Elastic LTB")
		fmt.Printf("Mn = %.0f k-ft\n", Mn_kft)
		return Mn
	} else { //case b (inelastic LTB)
		var Mn float64 = Cb * (Mp - (Mp-0.7*Fy*shape.Sx)*((Lb-Lp)/(Lr-Lp))) //k-in
		var Mn_kft float64 = Mn / 12                                        //k-ft
		fmt.Println("Case B: Limit State is Inelastic LTB")
		fmt.Printf("Mn = %.0f k-ft\n", Mn_kft)
		return Mn
	}
}

func flange_check(shape sections.WShape, Fy float64, E float64) string {
	if shape.Bf/shape.Tf < 0.38*math.Sqrt(E/Fy) {
		return "compact"
	} else if shape.Bf/shape.Tf > math.Sqrt(E/Fy) {
		return "slender"
	} else {
		return "noncompact"
	}
}

func web_check(shape sections.WShape, Fy float64, E float64) string {
	if shape.H/shape.Tw < 3.76*math.Sqrt(E/Fy) {
		return "compact"
	} else if shape.H/shape.Tw > 5.70*math.Sqrt(E/Fy) {
		return "slender"
	} else {
		return "noncompact"
	}
}

// Determine which sections apply:
func Beam_capacity(shape sections.WShape, Cb float64, Fy float64, E float64, Lb float64) float64 {
	//Check flange condition:
	var flange_condition string = flange_check(shape, Fy, E)
	fmt.Printf("Flange is %s\n", flange_condition)

	//Check web condition:
	var web_condition string = web_check(shape, Fy, E)
	fmt.Printf("Web is %s\n", web_condition)

	// Always calc F2
	var Mn float64 = Mn_Calc(shape, Cb, Fy, E, Lb)
	return Mn

	//Compact Web & Noncompact Flanges:
	//F2 applies
	//F3-1 applies

	//Compact Web & Slender Flanges:
	//F2 applies
	//F3-2 applies

	//Noncompact Web:
	//F4 applies

	//Slender flanges:

}
