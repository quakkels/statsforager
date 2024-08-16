# Docker Notes

## Run PostrgreSQL in Docker

Example command for running a docker container with PostgreSQL inside.

The `-e` environmnet variables are things that the `postgres` docker 
image know about.

```
docker run -p 5432:5432 -d \
    -e POSTGRES_PASSWORD=postgres \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_DB=stats \
    -v pgdata:/var/lib/postgresql/data \
    postgres
```

Once this command has been executed, docker will run the container in the 
background.

To connect `psql` or other tools to the `stats` database in the container, we
may use a command like this:

```
psql stats -h localhost -U postgres
```

The above command will use our natively running PostgreSQL tools, such as 
psql, to connect to the database.

If we didn't want to install postgres tools on our system, we can also connect 
to the container and use its `psql` to connect to the database.

```
docker exec -it {{IdHashOfContainer}} psql -U postgres stats
```

Breaking down the command, `exec` is short for execute, and `-it` means 
"interactive console`. It's followed by the ID of the container we want to 
connect to, which can be found with the command `docker container ls`. The 
rest is a regular `psql` command.
