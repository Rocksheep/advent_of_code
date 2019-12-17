package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type chemical struct {
	quantity int
	name     string
}

var leftOvers map[chemical]int = map[chemical]int{}

func main() {
	formulas := map[chemical][]chemical{}
	createFormulasFromInputFile(formulas)

	fuel := findChemicalFormula("FUEL", formulas)
	numOres := findOreRequirements(*fuel, formulas)

	fmt.Println(numOres)
}

func findOreRequirements(chem chemical, formulas map[chemical][]chemical) int {
	numOres := 0
	for _, chem := range formulas[chem] {
		if chem.name == "ORE" {
			numOres += chem.quantity
			continue
		}
		formula := findChemicalFormula(chem.name, formulas)

		for leftOvers[*formula] < chem.quantity {
			numOres += findOreRequirements(*formula, formulas)

			leftOvers[*formula] += formula.quantity
		}

		leftOvers[*formula] -= chem.quantity
	}

	return numOres
}

func createFormulasFromInputFile(formulas map[chemical][]chemical) {
	inputFile, err := os.Open("./input")
	reader := bufio.NewReader(inputFile)
	if err != nil {
		fmt.Println("Input not found")
		return
	}

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("ERROR", err)
			break
		}
		s := strings.Split(line, " => ")
		input := s[0]
		result := s[1]

		resultingChemical := chemical{}
		fmt.Sscanf(result, "%d %s", &resultingChemical.quantity, &resultingChemical.name)
		formulas[resultingChemical] = []chemical{}

		ingredients := strings.Split(input, ", ")
		for _, ingredient := range ingredients {
			part := chemical{}
			fmt.Sscanf(ingredient, "%d %s", &part.quantity, &part.name)
			formulas[resultingChemical] = append(formulas[resultingChemical], part)
		}
	}
}

func findChemicalFormula(name string, formulas map[chemical][]chemical) *chemical {
	for chem := range formulas {
		if chem.name == name {
			return &chem
		}
	}

	return nil
}
