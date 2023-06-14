package generators

import (
	"github.com/ortizdavid/dbsample/helpers"
	"github.com/ortizdavid/filemanager/core"
	"github.com/ortizdavid/message-helper/messages"
)

type ProcedureGenerator struct {
}

func (proc ProcedureGenerator) Generate(sampleName string, rdb string) {
	var argument *helpers.Argument
	fileManager := core.FileManager{}
	destFolder := sampleName+"-"+rdb
	file := sampleName+".sql"

	if argument.ContainsSample(argument.GetSamples(), sampleName, "procedure") == true {
		fileManager.CreateSingleFolder(destFolder)
		fileManager.CreateSingleFile(destFolder, file)
		messages.Success("Procedure '" + sampleName + "' Generated Successfully!")
	} else {
		messages.Error("Procedure '"+sampleName+"' does not exists")
	}
}