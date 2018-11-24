# try to read text file to mysql database
`mysql require host port 9999`

`mysql data volume mapping to .data in current directory`

## Installation mysql with sql script for create database and table
```sh
$   docker-compose up -d
$   docker exec -it txt-file-db bash
$   mysql -u root -p < /sql/script.sql
```
