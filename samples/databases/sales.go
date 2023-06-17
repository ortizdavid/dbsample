package samples

func (db *DatabaseSample) GetSalesSample(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = db.createDatabase("db_sales", "mysql") + db.createSalesTablesMySQL() + db.insertsSales()
	case "postgres":
		sql = db.createDatabase("db_sales", "postgres") + db.createSalesTablesPostgreSQL() + db.insertsSales()
	}
	return sql
}


func (db *DatabaseSample) createSalesTablesMySQL() string {
return `

CREATE TABLE continents (

);

CREATE TABLE Sales (

);

CREATE TABLE cities (

);

`
}


func (db *DatabaseSample) createSalesTablesPostgreSQL() string {
return `

CREATE TABLE continents (

);

CREATE TABLE Sales (

);

CREATE TABLE cities (

);

`
}


func (db *DatabaseSample) insertsSales() string {
return `

INSERT INTO continents (continent_name) VALUES ('Africa');

`
}