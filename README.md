# DBSample

Tool for generate database Samples

## Principal Features
- Generate database templates
- Generate normalizated database templates
- Generate Stored Procedures, incluinng transactions
- Generate SQL Views
- Generate Triggers

## GENERATING DATABASES:
- dbsample -sample db-countries -type database -rdb postgres
- dbsample -sample db-people -type database -rdb mysql
- dbsample -sample db-sales -type database -rdb mysql
- dbsample -sample db-user-roles -type database -rdb mysql

## GENERATING VIEWS
- dbsample -sample trg-lock-insert -type trigger -rdb postgres
- dbsample -sample trg-min -type trigger -rdb mysql


## GENERATING STORED PROCEDURES
- dbsample -sample sp-product-stock -type procedure -rdb postgres
- dbsample -sample sp-min -type procedure -rdb mysql


## GENERATING TRIGGERS
- dbsample -sample trg-lock-insert -type trigger -rdb postgres
- dbsample -sample trg-min -type trigger -rdb mysql