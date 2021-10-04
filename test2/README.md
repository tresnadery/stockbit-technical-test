# Endpoint Documentation
You can see the Endpoint documentation in [Endpoint Documentation](API.md)

# Manual Installation

## Requirement

* golang installed

* grpc installed
see [grpc documentation](https://grpc.io/docs/languages/go/quickstart/)

* mysql database

# Run Application

rename .env.example to .env and setup config with your local configuration

run go mod tidy on root of folder test2 : 
``` go mod tidy ```

run the application : 
``` go run main.go ```

# Migration Database

migrate file logs.sql on database/ddl folder to your local database