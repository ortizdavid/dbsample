package samples


func (trg *TriggerSample) GetTriggerCompleteSample(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = trg.createTrigger("view_complete", "mysql") + trg.createTriggerCompleteMySQL()

	case "postgres":
		sql = trg.createTrigger("view_complete", "postgres") + trg.createTriggerCompletePostgreSQL() 
	}
	return sql
}


func (trg *TriggerSample) createTriggerCompleteMySQL() string {
	return `

`
}

func (trg *TriggerSample) createTriggerCompletePostgreSQL() string {
	return `

`
}

