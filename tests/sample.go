package test

import "github.com/ortizdavid/dbsample/generators"

type sampleTest struct {
	name string
	sType string
	rdb  string
}

var sampleGenerator *generators.SampleGenerator