# Try to insert mysql database with data from text file
`mysql require host port 9999`

`mysql data volume mapping to .data in current directory`

**`mysql allow empty root password`**

`you can set mysql root password via $MYSQL_ROOT_PASSWORD environment variable`

## Installation mysql with sql script for create database and table
```sh
$   docker-compose up -d
$   docker exec -it txt-file-db bash
[inside container]$   mysql -u root -p < /sql/script.sql
```

## Download mysql driver
```sh
$   GOPATH=$PWD go get -u github.com/go-sql-driver/mysql
```

## How to run
```sh
$   GOPATH=$PWD go run main.go
```