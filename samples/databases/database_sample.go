package samples

type DatabaseSample struct {
}


func (db *DatabaseSample) createDatabase(dbName string, rdb string) string {
strDB := ""
switch rdb {
case "mysql":
strDB += `
--- DATABASE: db_user_roles
--- RDBMS MySQL---

CREATE DATABASE `+dbName+`
USE `+dbName+`;

`
case "postgres":
strDB += `
--- DATABASE: db_user_roles
--- RDBMS PostgreSQL---

CREATE DATABASE `+dbName+`
\c `+dbName+`;
`
}
return strDB
}