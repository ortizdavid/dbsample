package generators

import (
	"github.com/ortizdavid/dbsample/helpers"
	"github.com/ortizdavid/message-helper/messages"
)

type SampleGenerator struct {
}


func (generator *SampleGenerator) Generate(sampleName string, sampleType string, relationalDb string) bool {
	var argument *helpers.Argument
	
	if argument.Contains(argument.GetSamples(), sampleName) == false {
		messages.Error("Sample '"+sampleName+"' does not exists!")
		helpers.PrintSamples()
		return false
	} else if argument.Contains(argument.GetSampleTypes(), sampleType) == false {
		messages.Error("Sample Type '"+sampleType+"' does not exists!")
		helpers.PrintSampleTypes()
		return false
	} else if argument.Contains(argument.GetRelationalDBs(), relationalDb) == false {
		messages.Error("Relational Database '"+sampleName+"' does not exists!")
		helpers.PrintRelationalDBs()
		return false
	}  else {
		switch sampleType {
		case "database":
		    DatabaseGenerator{}.Generate(sampleName, relationalDb)
		case "view":
			ViewGenerator{}.Generate(sampleName, relationalDb)
		case "procedure":
			ProcedureGenerator{}.Generate(sampleName, relationalDb)
		case "trigger":
			TriggerGenerator{}.Generate(sampleName, relationalDb)
		}
		return true
	}
}