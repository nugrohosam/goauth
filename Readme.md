## Set project

- copy .env.example to .env.yaml
- configure server environment that you needs in file .env.yaml

## System Requirement protoc

- install go to you computer
- install protoc to your computer (optional) 

## Get All Modules

- $ go mod tidy
- $ go get

## Run Migration & Run Server

- $ go run migration.go
- $ go run main.go

## Note

if you want to create some migration, add files to migration/{version that you want}

- {number}{nameofmigration}.up.sql
- {number}{nameofmigration}.down.sql

## Checklist

- [x] Register Users
- [x] Auth JWT Login
- [x] Middleware access denied
- [x] Test api
- [x] Cleaner DB Test
- [x] Request gRPC
- [x] Test gRPC

