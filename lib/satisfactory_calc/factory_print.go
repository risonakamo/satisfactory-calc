// factory printing related funcs

package satisfactory_calc

import (
	"fmt"
	"math"
	"satisfactory-calc/lib/utils"

	"github.com/fatih/color"
)

// return factory formatted string.
// give min length to expand the resulting string to the size
func printFactoryStr(
    fac Factory,
    minLength int,
    indentLevel int,
    indentSize int,
) string {
    var indentStr string=utils.DuplicateString(" ",indentLevel*indentSize)

    var initial string=fmt.Sprintf("[ ] %s%s: %dx %s (%s) @ %s -> %dx%s = %s",
        indentStr,
        color.RedString("%d",indentLevel+1),
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
        spacingAdd=0
    }

    return fmt.Sprintf("[ ] %s%s: %dx %s (%s)%s@ %s -> %dx%s = %s",
        indentStr,
        color.RedString("%d",indentLevel+1),
        fac.BuilderCount,
        fac.ItemName,
        fac.RecipeName,
        utils.DuplicateString(" ",spacingAdd+4),
        color.YellowString("%.2f",fac.ClockRate),
        fac.BuilderCount,
        color.CyanString("%.2f",fac.OutputPerBuilder),
        color.HiGreenString("%.2f",fac.TotalOutput),
    )
}

// long print factory with initialiser options to recursive version
func longPrintFactory(fac Factory) {
    var result []string=longPrintFactory2(fac,0,0,4)

    var longestLen int=0
    var facStr string
    for _,facStr = range result {
        longestLen=int(math.Max(float64(longestLen),float64(len(facStr))))
    }

    result=longPrintFactory2(fac,longestLen,0,4)
    for _,facStr = range result {
        fmt.Println(facStr)
    }
}

// long print factory internal recursive version. returns list of all strings
// to be printed for a factory
func longPrintFactory2(
    fac Factory,
    minLength int,
    indentLevel int,
    indentSize int,
) []string {
    var result []string

    result=append(result,printFactoryStr(fac,minLength,indentLevel,indentSize))

    var subFacDict FactorybyRecipe
    for _,subFacDict = range fac.SubFactories {
        var subFac Factory=utils.GetDictFirstItem(subFacDict)

        result=append(result,longPrintFactory2(
            subFac,
            minLength,
            indentLevel+1,
            indentSize,
        )...)
    }

    return result
}