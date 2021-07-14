package main

import (
	"fmt"

	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

func main() {
	regex := pcre.MustCompile(`(\w+):\1`, 0)

	matches := regex.MatcherString("John:John", 0).Matches()
	fmt.Println("For John:John: ", matches)

	matches = regex.MatcherString("The names are Simon:Simon", 0).Matches()
	fmt.Println("For The names are Simon:Simon: ", matches)

	matches = regex.MatcherString("John:Simon", 0).Matches()
	fmt.Println("For John:Simon: ", matches)

}
