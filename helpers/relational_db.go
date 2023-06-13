package helpers

type RelationalDb struct {
	Name string
	Description string
}

func (rdb RelationalDb) GetArguments() []RelationalDb{
	return []RelationalDb {
		{ Name: "mysql", Description: "MySQL" },
		{ Name: "postgres", Description: "PostgreSQL" },
	}
}

