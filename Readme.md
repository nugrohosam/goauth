## Set project

- copy .env.example to .env.yaml
- configure server environment that you needs in file .env.yaml

## System Requirement

- install go to you computer
- install protoc to your computer (optional)

## Using Protocol Buff (following this if you are do install)

- copy folder `google`, from inside of folder `root_protoc/include` to `root_folder_this_project/third_party`
- make your file `.proto` as you need in `proto` folder
- generate go protobuff with command as noted in file `generateProto.txt`

## Get All Modules

- $ go mod tidy
- $ go get

## Run Migration & Run Server

- $ go run migration.go
- $ go run main.go

## Note

if you want to create some migration, run `$ migrate-create -namefile=namefileyouwant` it will automaticaly create 2 files in migrations folder 

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
- [x] Migration tools

