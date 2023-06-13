package helpers

type Sample struct {
	Name string
	Description string
}

func (sa Sample) GetSamples() []Sample {
	return []Sample {
		{ Name: "db-user-roles", Description: "Database with Users and Roles" },
		{ Name: "db-people", Description: "Database People Tables Structure" },
		{ Name: "db-countries", Description: "All Countries and Cities" },
		{ Name: "procedure", Description: "Procedure Sample, including transactions" },
		{ Name: "view", Description: "SQL View Sample" },
		{ Name: "trigger", Description: "Trigger Sample" },
	}
}