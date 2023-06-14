package samples


func (view *ViewSample) GetViewCompleteSample(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = view.createView("view_complete", "mysql") + view.createViewCompleteMySQL()

	case "postgres":
		sql = view.createView("view_complete", "postgres") + view.createViewCompletePostgreSQL() 
	}
	return sql
}


func (view *ViewSample) createViewCompleteMySQL() string {
	return `

`
}


func (view *ViewSample) createViewCompletePostgreSQL() string {
	return `

`
}

