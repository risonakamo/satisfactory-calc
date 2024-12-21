// run with go run (for now)

package main

import (
	"errors"
	"fmt"
	"maps"
	"os"
	"satisfactory-calc/lib/satisfactory_calc"
	"slices"

	"github.com/k0kubun/pp/v3"
	"github.com/manifoldco/promptui"
	"k8s.io/apimachinery/pkg/util/sets"
)

// program args
type CliArgs struct {
	ItemSelect string
	RecipeSelect string
}

func main() {
	var args CliArgs
	var e error
	args,e=getArgs()

	if e!=nil {
		fmt.Println("arguments warning:",e)
		fmt.Println("usage: {exe} {item name} {recipe name}")
	}

	var itemSelect string=args.ItemSelect
	var recipeSelect string=args.RecipeSelect

	var recipesData satisfactory_calc.RecipesDict=satisfactory_calc.LoadRecipesDict(
		"../../data/factorylab_data.json",
	)
	var allItemsRawResources sets.Set[string]=satisfactory_calc.
		RecipesDictToRawResourceSet(recipesData)

	e=checkItem(itemSelect,recipeSelect,recipesData)

	if e!=nil {
		return
	}

	var itemRecipe satisfactory_calc.ItemRecipe=recipesData[itemSelect][recipeSelect]

	var factory satisfactory_calc.Factory=satisfactory_calc.CreateFactory(itemRecipe)
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
				allItemsRawResources,
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

			var currentResources satisfactory_calc.InputsDict=
			satisfactory_calc.CalculateResourceUse(
				factory,
				sets.New[string](),
			)

			fmt.Println("Currently Needed Resources:")
			satisfactory_calc.PrintInputsDict(currentResources)
			pp.Println(factory)

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
	satisfactory_calc.PrintAlternatesDict(recipeErr.AvailableRecipes,recipeErr.NeededAmount)

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

// get program args
func getArgs() (CliArgs,error) {
	var args CliArgs=CliArgs{}

	if len(os.Args)==1 {
		return args,errors.New("not enough args")
	} else if len(os.Args)==2 {
		args.ItemSelect=os.Args[1]
		return args,errors.New("missing recipe name")
	}

	return CliArgs{
		ItemSelect: os.Args[1],
		RecipeSelect: os.Args[2],
	},nil
}

// check if item and recipe are in recipes dict. if not, prints helpful text and returns
// error
func checkItem(
	item string,
	recipe string,
	recipesData satisfactory_calc.RecipesDict,
) error {
	var in bool
	var foundRecipes satisfactory_calc.AlternatesDict
	foundRecipes,in=recipesData[item]

	if !in {
		fmt.Println("Could not find item:",item)
		fmt.Println("Available Items:")

		var item string
		for item = range recipesData {
			fmt.Println(item)
		}
		return errors.New("bad item")
	}

	_,in=foundRecipes[recipe]

	if !in {
		fmt.Println("Could not find recipe:",recipe)
		fmt.Println("Available Recipes:")
		var aRecipe string
		for aRecipe = range foundRecipes {
			fmt.Println(aRecipe)
		}
		return errors.New("bad recipe")
	}

	return nil
}