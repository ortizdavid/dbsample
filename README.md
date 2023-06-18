# DBSample

Tool for generate database Samples

## Principal Features
- Generate database templates
- Generate normalizated database templates
- Generate Stored Procedures, incluinng transactions
- Generate SQL Views
- Generate Triggers



## Generating Databases:
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