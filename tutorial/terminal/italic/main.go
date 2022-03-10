package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	whilte := color.New(color.FgWhite)
	boldWhite := whilte.Add(color.Italic)
	fmt.Println()
	boldWhite.Println("This will print text in italic white.")
	fmt.Println()
}
