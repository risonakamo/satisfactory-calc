// run with go run (for now)

package main

import (
	"fmt"
	"satisfactory-calc/lib/satisfactory_calc"
)

func main() {
	var itemSelect string="heavy-modular-frame"
	var recipeSelect string="Heavy Modular Frame"

	var recipesData satisfactory_calc.RecipesDict=satisfactory_calc.LoadRecipesDict(
		"../../data/factorylab_data.json",
	)

	var factory satisfactory_calc.Factory=satisfactory_calc.CreateFactory(itemSelect,recipeSelect)
	var e error
	var selectedRecipes []string

	for {
		factory,e=satisfactory_calc.ConstructFactory2(
			factory,
			recipesData,
			selectedRecipes,
		)

		if e==nil {
			fmt.Println("Successfully constructed factory")
			satisfactory_calc.LongPrintFactory(factory)
			// todo: print out total raw resources
			return
		}
	}
}