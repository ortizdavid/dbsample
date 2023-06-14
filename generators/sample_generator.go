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
		return false
	} else if argument.Contains(argument.GetSampleTypes(), sampleType) == false {
		messages.Error("Sample Type'"+sampleType+"' does not exists!")
		return false
	} else if argument.Contains(argument.GetRelationalDBs(), relationalDb) == false {
		messages.Error("Relational Database '"+sampleName+"' does not exists!")
		return false
	}  else {
		switch sampleType {
		case "database":
		    DatabaseGenerator{}.Generate(sampleName, relationalDb)
			break
		case "view":
			ViewGenerator{}.Generate(sampleName, relationalDb)
			break
		case "procedure":
			ProcedureGenerator{}.Generate(sampleName, relationalDb)
			break
		case "trigger":
			TriggerGenerator{}.Generate(sampleName, relationalDb)
			break
		}
		return true
	}
	
}