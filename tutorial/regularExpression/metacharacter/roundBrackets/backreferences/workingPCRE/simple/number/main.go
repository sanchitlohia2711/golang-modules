package main

import (
	"fmt"

	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
)

func main() {
	regex := pcre.MustCompile(`(\d)\1+`, 0)

	matches := regex.MatcherString("1111", 0).Matches()
	fmt.Println("For 1111 : ", matches)

	matches = regex.MatcherString("88888888", 0).Matches()
	fmt.Println("For 88888888 : ", matches)

	matches = regex.MatcherString("444", 0).Matches()
	fmt.Println("For 444 : ", matches)

	matches = regex.MatcherString("123", 0).Matches()
	fmt.Println("For 123 : ", matches)

}
