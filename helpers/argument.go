package helpers

import "fmt"

type Argument struct {
	Name        string
	Description string
	ArgType     string
}

func (argument *Argument) GetFlags() []Argument {
	return []Argument{
		{ Name: "-sample", Description: "Sample of Database, Trigger, Procedure, ..." },
		{ Name: "-type", Description: "Sample Type: Database, Procedure, View, Trigger" },
		{ Name: "-rdb", Description: "Relational Database {mysql, postgres}" },
		{ Name: "-help", Description: "Help for Usage" },
		{ Name: "-examples", Description: "Help for Examples" },
		{ Name: "-list-samples", Description: "List al Samples, including description" },
		{ Name: "-list-rdbs", Description: "List all Relational Databases" },
		{ Name: "-list-types", Description: "List all Relational Databases" },
	}
}

func (argument *Argument) GetSamples() []Argument {
	return []Argument{
		{ Name: "db-user-roles", ArgType: "database", Description: "Database with Users and Roles" },
		{ Name: "db-countries", ArgType: "database", Description: "All Countries and Cities" },
		{ Name: "db-people", ArgType: "database", Description: "Database People Tables Structure" },
		{ Name: "db-recruitment", ArgType: "database", Description: "Database for Recruitment System" },
		{ Name: "db-sales", ArgType: "database", Description: "Database for Sales" },
		{ Name: "view-product-data", ArgType: "view", Description: "SQL View Product normalizated data Sample" },
		{ Name: "view-min", ArgType: "view", Description: "Minimal SQL View Sample" },
		{ Name: "trg-lock-insert", ArgType: "trigger", Description: "Trigger to Lock Insert into a table" },
		{ Name: "trg-min", ArgType: "trigger", Description: "Minimal Trigger Sample" },
		{ Name: "sp-product-stock", ArgType: "procedure", Description: "Product Stock Procedure Sample, including transactions" },
		{ Name: "sp-min", ArgType: "procedure", Description: "Minimal Procedure Sample, including transactions" },
	}
}

func (argument *Argument) GetRelationalDBs() []Argument {
	return []Argument{
		{ Name: "mysql", Description: "MySQL" },
		{ Name: "postgres", Description: "PostgreSQL" },
	}
}

func (argument *Argument) GetSampleTypes() []Argument {
	return []Argument{
		{ Name: "database", Description: "Database" },
		{ Name: "procedure", Description: "Stored Procedure" },
		{ Name: "trigger", Description: "Trigger" },
		{ Name: "view", Description: "SQL View" },
	}
}

func (argument *Argument) Contains(args []Argument, name string) bool {
	for _, arg := range args {
		if name == arg.Name {
			return true
		}
	}
	return false
}

func (argument *Argument) ContainsSample(args []Argument, sampleName string, argType string) bool {
	for _, arg := range args {
		if sampleName == arg.Name && argType == arg.ArgType {
			return true
		}
	}
	return false
}

func (argument *Argument) PrintArguments(args []Argument) {
	for _, arg := range args {
		fmt.Printf("\t%s ----- %s\n", arg.Name, arg.Description)
	}
}