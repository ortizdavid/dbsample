# DBSample

Command Line Application to generate database templates/examples.

## Principal Features
- Generate database templates
- Generate normalizated database templates
- Generate Stored Procedures, incluinng transactions
- Generate SQL Views
- Generate Triggers


# Usage
dbsample -sample **<SAMPLE_NAME>** -type **<SAMPLE_TYPE>** -rdb **<RELATIONAL_DATABASE>**



## Samples
- **db-user-roles** -  Database with Users and Roles
- **db-countries** -  All Countries and Cities
- **db-people** -  Database People Tables Structure
- **db-recruitment** -  Database for Recruitment System
- **db-sales** -  Database for Sales
- **view-product-data** -  SQL View Product normalizated data Sample
- **view-min** -  Minimal SQL View Sample
- **trg-lock-insert** -  Trigger to Lock Insert into a table
- **trg-min** -  Minimal Trigger Sample
- **sp-product-stock** -  Product Stock Procedure Sample, including transactions
- **sp-min** -  Minimal Procedure Sample, including transactions



# Sample Types
- **database** - Database
- **procedure** - Stored Procedure
- **trigger** - Trigger
- **view** - SQL View



# Relational Databases
- **mysql** - MySQL        
- **postgres** - PostgreSQL



## Generating Databases
```console
dbsample -sample db-countries -type database -rdb postgres
```

```console
dbsample -sample db-people -type database -rdb mysql
```

```console
dbsample -sample db-sales -type database -rdb mysql
```

```console
dbsample -sample db-user-roles -type database -rdb mysql
```



## Gernerating Views
```console
dbsample -sample trg-lock-insert -type trigger -rdb postgres
```

```console
dbsample -sample trg-min -type trigger -rdb mysql
```



## Generating Stored Procedures
```console
dbsample -sample sp-product-stock -type procedure -rdb postgres
```

```console
dbsample -sample sp-min -type procedure -rdb mysql
```


## Generating Triggers
```console
dbsample -sample trg-lock-insert -type trigger -rdb postgres
```

```console
dbsample -sample trg-min -type trigger -rdb mysql
```


## Help Comand
```console
dbsample -help
```

## Examples
```console
dbsample -examples
```

## List Samples
```console
dbsample -list-samples
```

## List Sample types
```console
dbsample -list-types
```

## List Databases
```console
dbsample -list-rdbs
```


