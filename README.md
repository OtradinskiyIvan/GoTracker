# Go Tracker
Start db
```
docker run -d --name pg-tracker -e POSTGRES_PASSWORD=secret -p 5432:5432 postgres:16
```
Init db
```
docker exec -it pg-tracker psql -U postgres -c "CREATE DATABASE tracker;"
```