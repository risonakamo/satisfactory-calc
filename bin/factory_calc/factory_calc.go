package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"satisfactory-calc/lib/satisfactory_calc"
	"strings"

	"github.com/k0kubun/pp/v3"
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

	for {
		fmt.Println("Recipes:")
		pp.Println(selectedRecipes)

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

		if errors.As(e,&satisfactory_calc.MissingRecipeErrorE) {
			fmt.Println(e)
			fmt.Println("Enter a recipe to use:")
			fmt.Print("> ")

			var userSelect string

			var reader *bufio.Reader=bufio.NewReader(os.Stdin)
			userSelect,e=reader.ReadString('\n')

			if e!=nil {
				panic(e)
			}

			userSelect=strings.TrimSpace(userSelect)

			selectedRecipes=append(selectedRecipes,userSelect)
		} else {
			fmt.Println("unknown error")
			panic(e)
		}
	}
}