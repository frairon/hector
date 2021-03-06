package algo

import (
	"github.com/frairon/hector/core"
)

type Classifier interface {

	//Set training parameters from parameter map
	Init(params map[string]string)

	//Train model on a given dataset
	Train(dataset *core.DataSet)

	//Predict the probability of a sample to be positive sample
	Predict(sample *core.Sample) float64

	SaveModel(path string) error
	LoadModel(path string) error
}

type MultiClassClassifier interface {
	//Set training parameters from parameter map
	Init(params map[string]string)

	//Train model on a given dataset
	Train(dataset *core.DataSet)

	//Predict the probability of a sample to be positive sample
	PredictMultiClass(sample *core.Sample) *core.ArrayVector

	SaveModel(path string) error
	LoadModel(path string) error
}
