// components related to parsing factory lab json

package factorylab

// dict specifying amount of items required by an item.
// also used for outputs.
// key: item name
// val: number required
type InputsDict map[string]int

// defines the factory lab top level json file
// not fully complete - only includes the parts needed
// to be accessed
type FactorylabJson struct {
    Recipes []Recipe
}

// a recipe
type Recipe struct {
    // item string id. InputsDict use this
    Id string
    // display name of recipe
    Name string
    // seconds to produce N items (see output for number produced)
    Time int

    In InputsDict
    Out InputsDict
}