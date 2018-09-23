package main

import (
	"fmt"
	"log"
	"os"

	alg "github.com/ovidiumiron/stringCombination/findCombinations"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatalf("Wrong number arguments in command line:%v\n", args)
	}

	for s := range alg.FindCombinations(args[0]) {
		fmt.Printf("%s\n", s)
	}
}
