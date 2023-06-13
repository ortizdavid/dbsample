package generators

import "github.com/ortizdavid/message-helper/messages"

type ProcedureGenerator struct {
}

func (proc ProcedureGenerator) Generate(sampleName string, rdb string) {
	messages.Success("Procedure '"+sampleName+"' Generated Successfully!")
}