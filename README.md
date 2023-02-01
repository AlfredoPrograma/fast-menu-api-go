# FastMenu

**FastMenu** is a web application dedicated to majority of restaurants and food
companies, which want to generate _amazing_ digital menus.

## API

An API is exposed in order to handle operations over business resources of the
application.

## Setting up

### Prerequisites

- Go version >= 1.19
- Docker

### Installation

1. Run `go mod tidy` in order to install all _Go_ dependencies
2. Run copy the `.env.example` file into `.env` file and fullfill the key/value
   pairs
3. Run `docker-compose up -d` to run _PostgreSQL docker container_ in dettached
   mode
4. Run `go run ./cmd/main.go` to run the API in _"dev mode"_
5. Run `go build ./cmd/main.go` to compile application into binary file (For
   deploy or prod)

### _scripts Folder

There are utility `.sh` scripts inside `_scripts` directory:

- `serve-docs.sh`: Generates `swagger.json` file and serves the docs page
  _(needs to have [goswagger](https://goswagger.io/) installed to run it)_.

## Technologies

- Go:
  - [Echo](https://echo.labstack.com/) for http handling framework (v4.10.0)
  - [Zerolog](https://pkg.go.dev/github.com/rs/zerolog) for custom logger
    (v1.29.0)
  - [Viper](https://github.com/spf13/viper) for handling config and environment
    files (v1.15.0)
  - [Gorm](https://gorm.io/) as ORM (v1.24.3)
- Docker:
  - [Postgres](https://hub.docker.com/layers/library/postgres/15.1-alpine/images/sha256-69bd7cc493f84c970cf0af442dbf1ed15315b63c3a1cd2f8556dc9e98004a11f?context=explore)
    as main database management system (v15.1-alpine)
