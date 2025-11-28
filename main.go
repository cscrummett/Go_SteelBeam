package main

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/cscrummett/Go_SteelBeam/design"
	"github.com/cscrummett/Go_SteelBeam/sections"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "design-beam",
		Short: "Steel beam design calculator",
		Run: func(cmd *cobra.Command, args []string) {
			runBeamDesign()
		},
	}
	rootCmd.Execute()
}

func runBeamDesign() {
	// Survey prompts
	var shapeSelected string
	var Lb_ft float64
	var Cb float64 = 1.0  //conservative
	var Fy float64 = 50   //ksi
	var E float64 = 29000 //ksi

	// Load shape options from CSV
	shapeOptions, err := sections.LoadShapeOptions()
	if err != nil {
		fmt.Printf("Error loading shapes: %v\n", err)
		return
	}

	// Ask for beam shape
	shapePrompt := &survey.Select{
		Message: "Select beam shape:",
		Options: shapeOptions,
		Filter: func(filter string, value string, index int) bool {
			// Return true if value should be shown
			return strings.Contains(strings.ToUpper(value), strings.ToUpper(filter))
		}, //allows typing to filter (we like this typa thing!)
	}
	survey.AskOne(shapePrompt, &shapeSelected)

	// Ask for length
	lengthPrompt := &survey.Input{
		Message: "Enter beam length (ft):",
	}
	survey.AskOne(lengthPrompt, &Lb_ft)
	var Lb float64 = Lb_ft * 12

	shapes, _ := sections.LoadWShapes()
	var shape sections.WShape
	found := false
	for _, s := range shapes {
		if s.Shape == shapeSelected {
			shape = s
			found = true
			break
		}
	}

	if !found {
		panic("Shape not found: " + shapeSelected)
	}

	fmt.Println("")
	var Mn float64 = design.Mn_Calc(shape, Cb, Fy, E, Lb)
	var Mn_kft float64 = Mn / 12

	p := message.NewPrinter(language.English)
	p.Printf("\nPhiMn = %.0f k-ft\n", 0.9*Mn_kft)
}
