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

- $ go get
- $ go mod tidy

## Migration Create

if you want to create some migration, run `$ migrate-create -namefile=namefileyouwant` it will automaticaly create 2 files in migrations folder 

- {number}{nameofmigration}.up.sql
- {number}{nameofmigration}.down.sql

## Migration Up

if you want to create some migration, run `$ migrate -state=up -stage=dev` it will automaticaly migration up

## Migration Down

if you want to create some migration, run `$ migrate -state=down -stage=dev -howmuchstep=1` it will rollback 1 file .down.sql

## Run Server
- $ go run main.go

## Run Test API
- $ go test github.com/nugrohosam/gosampleapi/tests/api

## Run Test GRPC
- $ go test github.com/nugrohosam/gosampleapi/tests/grpc

## Checklist

- [x] Register Users
- [x] Auth JWT
- [x] Middlewares
- [x] Test api
- [x] Cleaner DB Test
- [x] Request gRPC
- [x] Test gRPC
- [x] Migration tools
- [x] Seeder tools

