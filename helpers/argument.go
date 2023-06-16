package helpers

type Argument struct {
	Name string
	Description string
	ArgType string
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

func (argument *Argument) GetFlags() []Argument {
	return []Argument {
		{ Name: "-sample", Description: "Sample of Database, Trigger, Procedure, ..." },
		{ Name: "-type", Description: "Sample Type: Database, Procedure, View, Trigger" },
		{ Name: "-rdb", Description: "Relational Database {mysql, postgres}" },
		{ Name: "-help", Description: "Help for Usage" },
		{ Name: "-list-samples", Description: "List al Samples, including description" },
		{ Name: "-list-rdbs", Description: "List all Relational Databases" },
		{ Name: "-list-types", Description: "List all Relational Databases" },
	} 
}

func (argument *Argument) GetSamples() []Argument {
	return []Argument {
		{ Name: "db-user-roles", ArgType: "database", Description: "Database with Users and Roles" },
		{ Name: "db-countries", ArgType: "database", Description: "All Countries and Cities" },
		{ Name: "db-people", ArgType: "database", Description: "Database People Tables Structure" },
		{ Name: "db-recruitment", ArgType: "database", Description: "Database for Recruitment System" },
		{ Name: "db-sales", ArgType: "database", Description: "Database for Sales" },
		{ Name: "view", ArgType: "view", Description: "SQL View Sample" },
		{ Name: "view-min", ArgType: "view", Description: "Minimal SQL View Sample" },
		{ Name: "trigger", ArgType: "trigger", Description: "Trigger Sample" },
		{ Name: "trigger-min", ArgType: "trigger", Description: "Minimal Trigger Sample" },
		{ Name: "procedure", ArgType: "procedure", Description: "Procedure Sample, including transactions" },
		{ Name: "procedure-min", ArgType: "procedure", Description: "Minimal Procedure Sample, including transactions" },
	}
}

func (argument *Argument) GetRelationalDBs() []Argument {
	return []Argument {
		{ Name: "mysql", Description: "MySQL" },
		{ Name: "postgres", Description: "PostgreSQL" },
	}
}

func (argument *Argument) GetSampleTypes() []Argument {
	return []Argument {
		{ Name: "database", Description: "Database" },
		{ Name: "procedure", Description: "Stored Procedure" },
		{ Name: "trigger", Description: "Trigger" },
		{ Name: "view", Description: "SQL View" },
	}
}

