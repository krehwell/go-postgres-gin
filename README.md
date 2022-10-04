# CRUD Operation with Gin and Postgres in Golang
Documentation which is made by me, for my needs, and is meant for my future dumb me who will forgot how did I made this.

## Setup DB (I'm using Arch Linux)
Installation
``` bash
# pacman -S postgres 
```

Initial DB Setup
``` bash
$ sudo -iu postgres
[postgres]$user initdb -D /var/lib/postgres/data # initialize default save for postgres
[postgres]$user createdb golangTestDB
```

You can now connect to DB, I'm using DBeaver to connect it, should be:
- File -> New -> Database Connection -> PostgreSQL Config it with
    - Server: `Host: localhost`, `Database: postgres`
    - Authentication: `Username: postgres`, `Password: <none>`. Leave password empty if never set before


## Setup Project
Installation
Clone the project then `go mod tidy` should install all the necessary deps

## Project Structure
```
|+ controllers  # logic handler for the API
|+ database     # db connection related, Postgres and GORM setup
|+ models       # structs
|+ routers      # end point by GIN
|  main.go      # main 😎
```
