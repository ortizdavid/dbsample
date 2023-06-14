package samples

type TriggerSample struct {
}


func (trg *TriggerSample) createTrigger(trgName string, rdb string) string {
return `
DROP TRIGGER IF EXISTS `+trgName+`;

CREATE TRIGGER `+trgName+`
`
}
