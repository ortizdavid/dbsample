package samples

func (db *DatabaseSample) GetShoppingCartSample(rdb string) string {
	sql := ""
	switch rdb {
	case "mysql":
		sql = db.createDatabase("db_shopping_cart", "mysql") + db.createShoppingCartTablesMySQL() + db.insertsShoppingCart()

	case "postgres":
		sql = db.createDatabase("db_shopping_cart", "postgres") + db.createShoppingCartTablesMySQL() + db.insertsShoppingCart()
	}
	return sql
}


func (db *DatabaseSample) createShoppingCartTablesMySQL() string {
	return `

CREATE TABLE continents (

);

CREATE TABLE ShoppingCart (

);

CREATE TABLE cities (

);

`
}


func (db *DatabaseSample) createShoppingCartTablesPostgreSQL() string {
	return `

CREATE TABLE continents (

);

CREATE TABLE ShoppingCart (

);

CREATE TABLE cities (

);

`
}


func (db *DatabaseSample) insertsShoppingCart() string {
	return `

INSERT INTO continents (continent_name) VALUES ('Africa');

`
}