package samples


func (proc *ProcedureSample) GetProcedureCompleteSample(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = proc.createProcedure("proc_complete", "mysql") + proc.createProcedureCompleteMySQL()

	case "postgres":
		sql = proc.createProcedure("proc_complete", "postgres") + proc.createProcedureCompletePostgreSQL() 
	}
	return sql
}


func (proc *ProcedureSample) createProcedureCompleteMySQL() string {
	return `

`
}


func (proc *ProcedureSample) createProcedureCompletePostgreSQL() string {
	return `

`
}

