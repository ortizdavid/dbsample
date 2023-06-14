package generators

import (
	"github.com/ortizdavid/dbsample/helpers"
	"github.com/ortizdavid/filemanager/core"
	"github.com/ortizdavid/message-helper/messages"
)

type TriggerGenerator struct {
}

func (trg TriggerGenerator) Generate(sampleName string, rdb string) {
	var argument *helpers.Argument
	fileManager := core.FileManager{}
	destFolder := sampleName+"-"+rdb
	file := sampleName+".sql"

	if argument.ContainsSample(argument.GetSamples(), sampleName, "trigger") == true {
		fileManager.CreateSingleFolder(destFolder)
		fileManager.CreateSingleFile(destFolder, file)
		messages.Success("Trigger '" + sampleName + "' Generated Successfully!")
	} else {
		messages.Error("Trigger '"+sampleName+"' does not exists")
	}
}