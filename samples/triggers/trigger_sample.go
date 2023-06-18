package samples

type TriggerSample struct {
}


func (trg *TriggerSample) createTrigger(trgName string, rdb string) string {
return `-- STORED PROCEDURE: `+trgName+`;
-- RELATIONAL DATABASE: `+rdb+`
`
}
