// components related to factory struct

package satisfactory_calc

// defines a single factory which builds 1 thing. factory consists
// of multiple builders
type Factory struct {
    // name of item. should correspond with an item name in
    // the recipes dict
    ItemName string

    // number of builders that make up this factory
    BuilderCount string
    // clock speed to set each builder in this factory to
    ClockRate float32

    // amount each builder unit in this factory should be making
    OutputPerBuilder float32
    // total output produced by this entire factory using all of its
    // builder units
    TotalOutput float32

    // requirements for each single builder in this factory.
    InputsPerBuilder InputsDict

    // total requirements of this factory. calculated by multiplying
    // inputs per builder by builder count.
    Inputs InputsDict

    // sub factories that feed into this factory
    SubFactories []Factory
}