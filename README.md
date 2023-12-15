# HTMX + Go to-do web app

### Populate environment variable(s)

```
vi ./postgres/.env
```

### Set up Docker container

```
docker build ./postgres/. -t postgres
```

```
docker run -p 5432:5432 --name postgres-server --env-file ./postgres/.env -d postgres
```

### Run application

```
go run .
```
