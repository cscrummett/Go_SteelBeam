package design

import (
	"fmt"
	"math"

	"github.com/cscrummett/Go_SteelBeam/sections"
)

func Mn_Calc(shape sections.WShape, Cb float64, Fy float64, E float64, Lb float64) float64 {
	var Lp float64 = 1.76 * shape.Ry * math.Sqrt(E/Fy)         //F2-5 (in) Limiting laterally unbraced length for yielding limit state
	var jcsxh0 float64 = shape.J * 1.0 / (shape.Sx * shape.H0) //c = 1.0 for W-flange (F2-8a)
	var rts float64 = math.Sqrt(math.Sqrt(shape.Iy*shape.Cw) / shape.Sx)

	var Mp float64 = Fy * shape.Zx // k-in
	var Mp_kft float64 = Mp / 12   //k-ft

	fmt.Println("Lb =", Lb/12, "ft")
	fmt.Println("Lp =", Lp/12, "ft")

	if Lb <= Lp { //case a (plastic hinge)
		fmt.Println("Case A: plastic hinge")
		fmt.Println("Fy =", Fy, " ksi")
		fmt.Println("Zx =", shape.Zx, "in^3")
		fmt.Println("Mp =", Mp_kft, "k-ft")
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
		fmt.Println("Case C: elastic LTB")
		fmt.Println("Mn =", Mn_kft, " k-ft")
		return Mn
	} else { //case b (inelastic LTB)
		var Mn float64 = Cb * (Mp - (Mp-0.7*Fy*shape.Sx)*((Lb-Lp)/(Lr-Lp))) //k-in
		var Mn_kft float64 = Mn / 12                                        //k-ft
		fmt.Println("Case B: inelastic LTB")
		fmt.Println("Mn =", Mn_kft, " k-ft")
		return Mn
	}
}
