package generators

import "github.com/ortizdavid/message-helper/messages"

type TriggerGenerator struct {
}

func (trg TriggerGenerator) Generate(sampleName string, rdb string) {
	messages.Success("Trigger '" + sampleName + "' Generated Successfully!")
}