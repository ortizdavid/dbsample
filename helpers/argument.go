package helpers

type Argument struct {
	Name string
	Description string
}

func (arg Sample) GetArguments() []Argument {
	return []Argument {
		{ Name: "-sample", Description: "Sample of Database, Trigger, Procedure, ..." },
		{ Name: "-rdb", Description: "Relational Database {mysql, postgres}" },
		{ Name: "-help", Description: "Help for Usage" },
		{ Name: "-list-samples", Description: "List al Samples, including description" },
		{ Name: "-list-rdbs", Description: "List all Relational Databases" },
	}
}