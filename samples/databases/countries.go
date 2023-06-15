package samples

func (db *DatabaseSample) GetCountriesSample(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = db.createDatabase("db_countries", "mysql") + db.createCountriesTablesMySQL() + db.insertsCountries()

	case "postgres":
		sql = db.createDatabase("db_countries", "postgres") + db.createCountriesTablesMySQL() + db.insertsCountries()
	}
	return sql
}


func (db *DatabaseSample) createCountriesTablesMySQL() string {
	return `

CREATE TABLE continents (

);

CREATE TABLE countries (

);

CREATE TABLE cities (

);

`
}


func (db *DatabaseSample) createCountriesTablesPostgreSQL() string {
	return `

CREATE TABLE continents (

);

CREATE TABLE countries (

);

CREATE TABLE cities (

);
`
}


func (db *DatabaseSample) insertsCountries() string {
	return `
INSERT INTO continents (continent_name) VALUES ('Africa');
`
}