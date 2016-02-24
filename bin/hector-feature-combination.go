package main

import (
	"github.com/frairon/hector"
	"github.com/frairon/hector/combine"
	"github.com/frairon/hector/core"
	"os"
	"strings"
)

func main() {
	train, _, _, _, params := hector.PrepareParams()

	feature_combination := combine.CategoryFeatureCombination{}
	feature_combination.Init(params)

	dataset := core.NewRawDataSet()
	dataset.Load(train)

	combinations := feature_combination.FindCombination(dataset)

	output := params["output"]

	file, _ := os.Create(output)
	defer file.Close()

	for _, combination := range combinations {
		file.WriteString(strings.Join(combination, "\t") + "\n")
	}
}
