# Run these commands in terminal


# GENERATING DATABASES:
- dbsample -sample db-countries -type database -rdb postgres
- dbsample -sample db-people -type database -rdb mysql
- dbsample -sample db-sales -type database -rdb mysql
- dbsample -sample db-user-roles -type database -rdb mysql


# GENERATING STORED PROCEDURES
- dbsample -sample sp-product-stock -type procedure -rdb postgres
- dbsample -sample sp-min -type procedure -rdb mysql


# GENERATING TRIGGERS
- dbsample -sample trg-lock-insert -type trigger -rdb postgres
- dbsample -sample trg-min -type trigger -rdb mysql

# GENERATING VIEWS
- dbsample -sample view-user-data -type trigger -rdb postgres
- dbsample -sample view-min -type trigger -rdb mysql