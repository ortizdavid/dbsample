package samples


func (proc *ProcedureSample) GetProcedureMinimalSample(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = proc.createProcedure("proc_minimal", "mysql") + proc.createProcedureMinimalMySQL()

	case "postgres":
		sql = proc.createProcedure("proc_minimal", "postgres") + proc.createProcedureMinimalPostgreSQL() 
	}
	return sql
}


func (proc *ProcedureSample) createProcedureMinimalMySQL() string {
return `

`
}


func (proc *ProcedureSample) createProcedureMinimalPostgreSQL() string {
return `

`
}

