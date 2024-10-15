# Golang REST API

Library REST API made with Go v1.23.2

### Users API endpoints:

- `GET /users/list`
- `POST /users/add`

### Borrowed books API endpoints:

- `GET /borrowed/{userId}`
- `GET /returned/{userId}`
- `POST /borrowed/add`
- `PUT /returned/update`

### Books API endpoints:

- `GET /books/list`

## Database

Create PostgreSQL database

`CREATE DATABASE library ENCODING UTF-8;`

## Project setup

Cd into project folder

`cd github.com/saks07/go-api`

In current folder create .env file

`cp .env.example .env`

and update Database credentials

Install [Golang Migrate](https://github.com/golang-migrate/migrate)

Setup Go dependencies

`go install`

Run Go project

`go run main.go`

After executing the command, Migrations will run, which will setup your database tables and mock data (code in main.go, main() method)
