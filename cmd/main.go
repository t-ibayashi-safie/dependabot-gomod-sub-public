package main

import (
	"github.com/t-ibayashi-safie/dependabot-gomod-sub-public/pkg/variables"
)

func main() {
	subVal := variables.GetVariables()
	println("subVal: " + subVal)
}
