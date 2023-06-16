package generators

import (
	"github.com/ortizdavid/dbsample/helpers"
	samples "github.com/ortizdavid/dbsample/samples/procedures"
	"github.com/ortizdavid/filemanager/core"
	"github.com/ortizdavid/message-helper/messages"
)

type ProcedureGenerator struct {
}

func (proc ProcedureGenerator) Generate(sampleName string, rdb string) {
	var argument *helpers.Argument
	var sample *samples.ProcedureSample
	fileManager := core.FileManager{}

	if argument.ContainsSample(argument.GetSamples(), sampleName, "procedure") == true {
		destFolder := sampleName+"-"+rdb
		file := sampleName+".sql"
		contentFile := ""
		fileManager.CreateSingleFolder(destFolder)
		fileManager.CreateSingleFile(destFolder, file)
		switch sampleName {
		case "sproc":
			contentFile = sample.GetProcedureCompleteSample(rdb)
		case "sproc-min":
			contentFile = sample.GetProcedureMinimalSample(rdb)
		}
		fileManager.WriteFile(destFolder, file, contentFile)
		messages.Success("Procedure '" + sampleName + "' Generated Successfully!")
	} else {
		messages.Error("Procedure '"+sampleName+"' does not exists")
	}
}