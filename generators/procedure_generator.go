package generators

import (
	"github.com/ortizdavid/dbsample/helpers"
	samples "github.com/ortizdavid/dbsample/samples/procedures"
	"github.com/ortizdavid/go-nopain/filemanager"
	"github.com/ortizdavid/go-nopain/messages"
)

type ProcedureGenerator struct {
}

func (proc ProcedureGenerator) Generate(sampleName string, rdb string) {
	var argument *helpers.Argument
	var sample *samples.ProcedureSample
	var fileManager filemanager.FileManager

	if argument.ContainsSample(argument.GetSamples(), sampleName, "procedure") == true {
		destFolder := sampleName+"-"+rdb
		file := sampleName+".sql"
		contentFile := ""
		fileManager.CreateSingleFolder(destFolder)
		fileManager.CreateSingleFile(destFolder, file)
		switch sampleName {
		case "sp-product-stock":
			contentFile = sample.GetProcedureProductStockSample(rdb)
		case "sp-min":
			contentFile = sample.GetProcedureMinimalSample(rdb)
		}
		fileManager.WriteFile(destFolder, file, contentFile)
		messages.Success("Procedure '" + sampleName + "' Generated Successfully!")
	} else {
		messages.Error("Procedure '"+sampleName+"' does not exists")
	}
}