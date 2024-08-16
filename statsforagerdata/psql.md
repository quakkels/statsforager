# `psql` Tips

## Start `psql` connected to a database
`psql {{dbname}} -h localhost -U postgres`

## List databases

- `\l`

## Change current database

- `\connect DBNAME`
- `\c DBNAME`

## List tables of current database

- `\dt`

## List columns of specific table

- `\d+ TABLENAME`
    - Includes columns, types, indexes, and references
- `SELECT column_name, * FROM information_schema.columns WHERE table_catalog = 'your_db_name' AND table_name = 'your_table';


