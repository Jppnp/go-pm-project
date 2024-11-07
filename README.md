# go-pm-project

backend api for kanban board(CRUD)

**Gin Framework
Postgres database server
Database versioning with goose**

# How to create new database version

`goose -dir mirgrations create <your_file_name> sql`

# How to run the service

### Prerequisite

* [X] golang version 1.22.x ++
* [X] postgres server
* [X] goose CLI

### Steps

* First thing first you need to run your postgres server because the service will migrate database first(if it has a new file) then run the service
* Then run type `go run main.go` to run main.go
