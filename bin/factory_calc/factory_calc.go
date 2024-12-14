// run with go run (for now)

package main

import (
	"errors"
	"fmt"
	"maps"
	"satisfactory-calc/lib/satisfactory_calc"
	"slices"

	"github.com/manifoldco/promptui"
)

func main() {
	var itemSelect string="heavy-modular-frame"
	var recipeSelect string="Heavy Modular Frame"

	var recipesData satisfactory_calc.RecipesDict=satisfactory_calc.LoadRecipesDict(
		"../../data/factorylab_data.json",
	)

	var itemRecipe satisfactory_calc.ItemRecipe=recipesData[itemSelect][recipeSelect]

	var factory satisfactory_calc.Factory=satisfactory_calc.CreateFactory(itemRecipe)
	var e error
	var selectedRecipes []string
	var constructCount int=0

	for {
		constructCount++
		fmt.Printf("Constructing... (%d)\n",constructCount)
		factory,e=satisfactory_calc.ConstructFactory2(
			factory,
			recipesData,
			selectedRecipes,
		)



		// completion operations
		if e==nil {
			fmt.Println()
			fmt.Println("Successfully constructed factory:")
			satisfactory_calc.LongPrintFactory(factory)
			fmt.Println()

			var resources satisfactory_calc.InputsDict=satisfactory_calc.CalculateResourceUse(
				factory,
				satisfactory_calc.DefaultRawResources,
			)

			fmt.Println("Total Resources:")
			satisfactory_calc.PrintInputsDict(resources)

			return
		}



		// missing recipe. have user select a recipe
		if errors.As(e,&satisfactory_calc.MissingRecipeErrorE) {
			fmt.Println()
			var recipeErr *satisfactory_calc.MissingRecipeError
			var ok bool
			recipeErr,ok=e.(*satisfactory_calc.MissingRecipeError)

			if !ok {
				panic("bad error cast")
			}

			selectedRecipes=append(selectedRecipes,userChooseRecipe(*recipeErr))



		// random unknown error occured
		} else {
			fmt.Println("unknown error")
			panic(e)
		}
	}
}

// given a missing recipe error, prompt user to select an available recipe.
// return the selected recipe.
func userChooseRecipe(recipeErr satisfactory_calc.MissingRecipeError) string {
	fmt.Println()
	fmt.Println()
	fmt.Println("-------------------------------")
	fmt.Println(recipeErr.Error())

	var choices []string=slices.Collect(maps.Keys(recipeErr.AvailableRecipes))

	var prompter promptui.Select=promptui.Select{
		Label: "Select Recipe",
		Items: choices,
		HideHelp: true,
	}

	var userSelect string
	var e error
	_,userSelect,e=prompter.Run()

	if e!=nil {
		panic(e)
	}

	return userSelect
}