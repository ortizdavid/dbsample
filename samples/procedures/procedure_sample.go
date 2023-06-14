package samples

type ProcedureSample struct {
}

func (proc *ProcedureSample) createProcedure(procName string, rdb string) string {
return `

DROP PROCEDURE IF EXISTS `+procName+`;

CREATE PROCEDURE `+procName+`

`
}