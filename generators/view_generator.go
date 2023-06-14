package generators

import (
	"github.com/ortizdavid/dbsample/helpers"
	"github.com/ortizdavid/filemanager/core"
	"github.com/ortizdavid/message-helper/messages"
)

type ViewGenerator struct {
}

func (view ViewGenerator) Generate(sampleName string, rdb string) {
	var argument *helpers.Argument
	fileManager := core.FileManager{}
	destFolder := sampleName+"-"+rdb
	file := sampleName+".sql"

	if argument.ContainsSample(argument.GetSamples(), sampleName, "view") == true {
		fileManager.CreateSingleFolder(destFolder)
		fileManager.CreateSingleFile(destFolder, file)
		messages.Success("View '" + sampleName + "' Generated Successfully!")
	} else {
		messages.Error("View '"+sampleName+"' does not exists")
	}
}