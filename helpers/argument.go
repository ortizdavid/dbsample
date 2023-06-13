package helpers

type Argument struct {
	Name string
	Description string
}

func (arg *Argument) Contains(args []Argument, name string) bool  {
	for _, arg := range args {
		if name == arg.Name {
			return true
		}
	}
	return false
}

func (arg *Argument) GetFlags() []Argument {
	return []Argument {
		{ Name: "-sample", Description: "Sample of Database, Trigger, Procedure, ..." },
		{ Name: "-type", Description: "Sample Type: Database, Procedure, View, Trigger" },
		{ Name: "-rdb", Description: "Relational Database {mysql, postgres}" },
		{ Name: "-help", Description: "Help for Usage" },
		{ Name: "-list-samples", Description: "List al Samples, including description" },
		{ Name: "-list-rdbs", Description: "List all Relational Databases" },
	} 
}

func (arg *Argument) GetSamples() []Argument {
	return []Argument {
		{ Name: "db-user-roles", Description: "Database with Users and Roles" },
		{ Name: "db-people", Description: "Database People Tables Structure" },
		{ Name: "db-countries", Description: "All Countries and Cities" },
		{ Name: "db-shopping-cart", Description: "All Countries and Cities" },
		{ Name: "view", Description: "SQL View Sample" },
		{ Name: "view-min", Description: "Minimal SQL View Sample" },
		{ Name: "trigger", Description: "Trigger Sample" },
		{ Name: "trigger-min", Description: "Minimal Trigger Sample" },
		{ Name: "procedure", Description: "Procedure Sample, including transactions" },
		{ Name: "procedure-min", Description: "Minimal Procedure Sample, including transactions" },
	}
}

func (arg *Argument) GetRelationalDBs() []Argument {
	return []Argument {
		{ Name: "mysql", Description: "MySQL" },
		{ Name: "postgres", Description: "PostgreSQL" },
	}
}

func (arg *Argument) GetSampleTypes() []Argument {
	return []Argument {
		{ Name: "database", Description: "Database" },
		{ Name: "procedure", Description: "Stored Procedure" },
		{ Name: "trigger", Description: "Trigger" },
		{ Name: "view", Description: "SQL View" },
	}
}

