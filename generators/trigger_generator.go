package generators

import (
	"github.com/ortizdavid/dbsample/helpers"
	samples "github.com/ortizdavid/dbsample/samples/triggers"
	"github.com/ortizdavid/go-nopain/filemanager"
	"github.com/ortizdavid/go-nopain/messages"
)

type TriggerGenerator struct {
}

func (trg TriggerGenerator) Generate(sampleName string, rdb string) {
	var argument *helpers.Argument
	var sample *samples.TriggerSample
	var fileManager filemanager.FileManager
	
	if argument.ContainsSample(argument.GetSamples(), sampleName, "trigger") == true {
		destFolder := sampleName+"-"+rdb
		file := sampleName+".sql"
		contentFile := ""
		fileManager.CreateSingleFolder(destFolder)
		fileManager.CreateSingleFile(destFolder, file)
		switch sampleName {
		case "trg-lock-insert":
			contentFile = sample.GetTriggerLockInsert(rdb)
		case "trg-min":
			contentFile = sample.GetTriggerMinimalSample(rdb)
		}
		fileManager.WriteFile(destFolder, file, contentFile)
		messages.Success("Trigger '" + sampleName + "' Generated Successfully!")
	} else {
		messages.Error("Trigger '"+sampleName+"' does not exists!")
	}
}