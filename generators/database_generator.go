package generators

import "github.com/ortizdavid/message-helper/messages"

type DatabaseGenerator struct {
}

func (db DatabaseGenerator) Generate(sampleName string, rdb string) {
	messages.Success("Database '" + sampleName + "' Generated Successfully!")
}