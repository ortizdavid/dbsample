package generators

import (
	"github.com/ortizdavid/dbsample/helpers"
	"github.com/ortizdavid/filemanager/core"
	"github.com/ortizdavid/message-helper/messages"
)

type DatabaseGenerator struct {
}

func (db DatabaseGenerator) Generate(sampleName string, rdb string) {
	var argument *helpers.Argument
	fileManager := core.FileManager{}
	destFolder := sampleName+"-"+rdb
	file := sampleName+".sql"

	if argument.ContainsSample(argument.GetSamples(), sampleName, "database") == true {
		fileManager.CreateSingleFolder(destFolder)
		fileManager.CreateSingleFile(destFolder, file)
		messages.Success("Database '" + sampleName + "' Generated Successfully!")
	} else {
		messages.Error("Database '"+sampleName+"' does not exists")
	}
}