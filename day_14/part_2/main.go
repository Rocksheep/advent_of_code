package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

type Recipe struct {
	chemical    Chemical
	ingredients []Chemical
}

type Chemical struct {
	quantity int
	name     string
}

func main() {
	recipes := createFormulasFromInputFile()
	excess := map[string]int{}

	numOres := 0
	var i int
	for i = 1; numOres < 1000000000000; i++ {
		numOres = findOreRequirements("FUEL", i, recipes, excess)
		fmt.Println("Fuel", i, "Ores", numOres)
	}

	fmt.Println(i - 2)
}

func findOreRequirements(chemicalName string, quantity int, recipes map[string]Recipe, excess map[string]int) int {
	if chemicalName == "ORE" {
		return quantity
	}

	if excess[chemicalName] >= quantity {
		excess[chemicalName] -= quantity
		return 0
	}

	if excess[chemicalName] > 0 {
		quantity -= excess[chemicalName]
		excess[chemicalName] = 0
	}

	recipe := recipes[chemicalName]
	batches := int(math.Ceil(float64(quantity) / float64(recipe.chemical.quantity)))

	ores := 0
	for _, ingredient := range recipe.ingredients {
		ores += findOreRequirements(ingredient.name, ingredient.quantity*batches, recipes, excess)
	}

	producedIngredients := batches * recipe.chemical.quantity
	excess[chemicalName] += producedIngredients - quantity

	return ores
}

func createFormulasFromInputFile() map[string]Recipe {
	recipes := map[string]Recipe{}

	inputFile, err := os.Open("./input")
	reader := bufio.NewReader(inputFile)
	if err != nil {
		fmt.Println("Input not found")
		return map[string]Recipe{}
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

		resultingChemical := Chemical{}
		fmt.Sscanf(result, "%d %s", &resultingChemical.quantity, &resultingChemical.name)
		recipe := Recipe{resultingChemical, []Chemical{}}

		ingredients := strings.Split(input, ", ")
		for _, ingredient := range ingredients {
			part := Chemical{}
			fmt.Sscanf(ingredient, "%d %s", &part.quantity, &part.name)
			recipe.ingredients = append(recipe.ingredients, part)
		}
		recipes[resultingChemical.name] = recipe
	}

	return recipes
}
