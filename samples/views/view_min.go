package samples


func (view *ViewSample) GetViewMinimalSample(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = view.createView("view_minimal", "mysql") + view.createViewMinimalMySQL()

	case "postgres":
		sql = view.createView("view_minimal", "postgres") + view.createViewMinimalPostgreSQL() 
	}
	return sql
}


func (view *ViewSample) createViewMinimalMySQL() string {
	return `

`
}


func (view *ViewSample) createViewMinimalPostgreSQL() string {
	return `

`
}

