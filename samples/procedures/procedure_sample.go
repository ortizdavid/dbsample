package samples

type ProcedureSample struct {
}

func (proc *ProcedureSample) createProcedure(procName string, rdb string) string {
return `-- STORED PROCEDURE: `+procName+`;

-- RELATIONAL DATABASE: `+rdb+`

`
}