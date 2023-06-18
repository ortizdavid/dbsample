package main

import (
	"os"
	"github.com/ortizdavid/dbsample/generators"
	"github.com/ortizdavid/dbsample/helpers"
)


func main() {

	cliArgs := os.Args
	numArgs := len(cliArgs)

	if numArgs <= 1 {
		helpers.PrintUsage()

	} else {
		
		if numArgs == 2 &&  cliArgs[2] == "-help" {
			helpers.PrintHelp()

		} else if numArgs == 2 &&  cliArgs[2] == "-examples" {
			helpers.PrintExamples()
	
		} else if numArgs == 2 &&  cliArgs[2] == "-list-samples" {
			helpers.PrintSamples()

		} else if numArgs == 2 &&  cliArgs[2] == "-list-types" {
			helpers.PrintSampleTypes()

		} else if numArgs == 2 &&  cliArgs[2] == "-list-rdbs" {
			helpers.PrintRelationalDBs()

		} else if numArgs == 7 {
			sampleName := cliArgs[2]
			sampleType := cliArgs[4]
			relationalDB := cliArgs[6]
			generator := generators.SampleGenerator{}
			generator.Generate(sampleName, sampleType, relationalDB)
			
		} else {
			helpers.PrintHelp()
		}
	}

}