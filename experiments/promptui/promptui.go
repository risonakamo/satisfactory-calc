// used with go run

package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {
    fmt.Println("some other prints")
    fmt.Println("some other prints")
    fmt.Println("some other prints")
    fmt.Println("some other prints")

    prompt:=promptui.Select{
        Label: "something",
        Items: []string{
            "a",
            "b",
            "c",
        },
        HideHelp: true,
    }

    _,result,e:=prompt.Run()

    if e!=nil {
        panic(e)
    }

    fmt.Println("selected",result)
}