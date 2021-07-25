## Try this project free
- postman docs : https://documenter.getpostman.com/view/4473147/TVzXDFka

## Set project

- copy .env.example to .env.yaml
- configure server environment that you needs in file .env.yaml

## System Requirement

- install go to you computer
- download & install protoc to your computer (optional) as your arc you use (windows/linux/macOS): download link https://github.com/protocolbuffers/protobuf/releases
## Using Protocol Buff (following this if you are do install)

- copy folder `google`, from inside of folder `root_protoc/include` to `root_folder_this_project/third_party`
- make your file `.proto` as you need in `assets/proto` folder
- generate go protobuff with command as noted in file `generateProto.txt`
- #makesure your GOPATH environment variable is set

## Get All Modules

- $ go get
- $ go mod tidy
## Run Makefile

- $ make

## Seeder

if you want to seder data, run `$ ./cmd/seeder -file=namefile.yaml`

## Run Server
- $ go run main.go --env-root-path=./ --service=[http, grpc]

## Run Test API
- $ go test github.com/nugrohosam/gosampleapi/tests/api

## Run Test GRPC
- $ go test github.com/nugrohosam/gosampleapi/tests/grpc

## Checklist

- [x] Supporting for database pgsql, mysql
- [x] Register Users
- [x] Auth JWT
- [x] Middlewares
- [x] Test api
- [x] DB Cleaner In Test
- [x] Request gRPC
- [x] Test gRPC
- [x] Migration tools
- [x] Session redis support
- [x] Pagination
- [x] Seeder tools
- [x] Kafka features
- [x] Encrypted token
- [x] Heroku deployment
- [x] CircleCI deployment
