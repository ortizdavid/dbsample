package samples

type ViewSample struct {
}

func (view *ViewSample) createView(viewName string, rdb string) string {
strView := `DROP VIEW IF EXISTS `+viewName+`;`
switch rdb {
case "mysql":
strView += `
CREATE VIEW `+viewName+` AS;

`
break
case "postgres":
strView += `
CREATE VIEW `+viewName+` AS;

`
break
}
return strView
}