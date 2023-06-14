package samples


func (trg* TriggerSample) GetTriggerMinimalSample(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = trg.createTrigger("view_min", "mysql") + trg.createTriggerMinimalMySQL()

	case "postgres":
		sql = trg.createTrigger("view_min", "postgres") + trg.createTriggerMinimalPostgreSQL() 
	}
	return sql
}


func (trg* TriggerSample) createTriggerMinimalMySQL() string {
	return `

`
}

func (trg*TriggerSample) createTriggerMinimalPostgreSQL() string {
	return `

`
}
