package generators

import (
	"github.com/ortizdavid/dbsample/helpers"
	samples "github.com/ortizdavid/dbsample/samples/views"
	"github.com/ortizdavid/go-nopain/filemanager"
	"github.com/ortizdavid/go-nopain/messages"
)

type ViewGenerator struct {
}

func (view ViewGenerator) Generate(sampleName string, rdb string) {
	var argument *helpers.Argument
	var sample *samples.ViewSample
	var fileManager filemanager.FileManager

	if argument.ContainsSample(argument.GetSamples(), sampleName, "view") == true {
		destFolder := sampleName+"-"+rdb
		file := sampleName+".sql"
		contentFile := ""
		fileManager.CreateSingleFolder(destFolder)
		fileManager.CreateSingleFile(destFolder, file)
		switch sampleName {
		case "view-product-data":
			contentFile = sample.GetViewProductDataSample(rdb)
		case "view-min":
			contentFile = sample.GetViewMinimalSample(rdb)
		}
		fileManager.WriteFile(destFolder, file, contentFile)
		messages.Success("View '" + sampleName + "' Generated Successfully!")
	} else {
		messages.Error("View '"+sampleName+"' does not exists!")
	}
}