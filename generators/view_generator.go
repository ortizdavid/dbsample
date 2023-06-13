package generators

import "github.com/ortizdavid/message-helper/messages"

type ViewGenerator struct {
}

func (view ViewGenerator) Generate(sampleName string, rdb string) {
	messages.Success("View '" + sampleName + "' Generated Successfully!")
}