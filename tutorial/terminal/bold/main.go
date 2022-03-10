package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	whilte := color.New(color.FgWhite)

	boldWhite := whilte.Add(color.Bold)
	fmt.Println()
	boldWhite.Println("Text in bold white.")
	fmt.Println()
}
