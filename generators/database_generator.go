package generators

import (
	"github.com/ortizdavid/dbsample/helpers"
	samples "github.com/ortizdavid/dbsample/samples/databases"
	"github.com/ortizdavid/filemanager/core"
	"github.com/ortizdavid/message-helper/messages"
)

type DatabaseGenerator struct {
}

func (db DatabaseGenerator) Generate(sampleName string, rdb string) {
	var argument *helpers.Argument
	var sample *samples.DatabaseSample
	fileManager := core.FileManager{}

	if argument.ContainsSample(argument.GetSamples(), sampleName, "database") == true {
		destFolder := sampleName+"-"+rdb
		file := sampleName+".sql"
		contentFile := ""
		fileManager.CreateSingleFolder(destFolder)
		fileManager.CreateSingleFile(destFolder, file)
		switch sampleName {
		case "db-user-roles":
			contentFile = sample.GetUserRolesSample(rdb)
		case "db-shopping-cart":
			contentFile = sample.GetShoppingCartSample(rdb)
		case "db-countries":
			contentFile = sample.GetCountriesSample(rdb)
		}
		fileManager.WriteFile(destFolder, file, contentFile)
		messages.Success("Database '" + sampleName + "' Generated Successfully!")
	} else {
		messages.Error("Database '"+sampleName+"' does not exists")
	}
}