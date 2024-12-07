// factory printing related funcs

package satisfactory_calc

import (
	"fmt"
	"satisfactory-calc/lib/utils"

	"github.com/fatih/color"
)

// format print single factory
func printFactory(
    fac Factory,
    minLength int,
) {
    fmt.Println(printFactoryStr(
        fac,
        minLength,
    ))
}

// return factory formatted string.
// give min length to expand the resulting string to the size
func printFactoryStr(
    fac Factory,
    minLength int,
) string {
    var initial string=fmt.Sprintf("%dx %s (%s) @ %s -> %dx%s = %s",
        fac.BuilderCount,
        fac.ItemName,
        fac.RecipeName,
        color.YellowString("%.2f",fac.ClockRate),
        fac.BuilderCount,
        color.CyanString("%.2f",fac.OutputPerBuilder),
        color.HiGreenString("%.2f",fac.TotalOutput),
    )

    var initialLen int=len(initial)

    var spacingAdd int=minLength-initialLen

    if spacingAdd<=0 {
        spacingAdd=1
    }

    return fmt.Sprintf("%dx %s (%s)%s@ %s -> %dx%s = %s",
        fac.BuilderCount,
        fac.ItemName,
        fac.RecipeName,
        utils.DuplicateString(" ",spacingAdd),
        color.YellowString("%.2f",fac.ClockRate),
        fac.BuilderCount,
        color.CyanString("%.2f",fac.OutputPerBuilder),
        color.HiGreenString("%.2f",fac.TotalOutput),
    )
}

// print factory and all subfactories
func longPrintFactory(
    fac Factory,
    indentLevel int,
    indentSize int,
    longestSize int,
) {
    var indentStr string=utils.DuplicateString(" ",indentLevel*indentSize)
    fmt.Printf("[ ] %s%s: ",
        indentStr,
        color.RedString("%d",indentLevel+1),
    )
    printFactory(fac,longestSize)

    var subFactoryDict FactorybyRecipe
    for _,subFactoryDict = range fac.SubFactories {
        var subFactory Factory=utils.GetDictFirstItem(subFactoryDict)
        longPrintFactory(subFactory,indentLevel+1,indentSize,longestSize)
    }
}