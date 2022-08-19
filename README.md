# simple-transaction-api

## Installation

Make sure you already installed go in your device, and all packages of Go language.

First you should clone this repository by ssh in your local device


```sh
git clone https://github.com/apryono/simple-transaction-api
```

After that, open this project and please run this command :
```sh
go mod tidy
```
This is mean to get or download all libraries in this project.

## Setup 

Then, before you run this project, please create file .env
```sh
mkdir .env
```
or you can change name of .env-example to .env and filled in according to your database. 

And filled your environment with this sample .env, this sample available in this project as .env-example, and this sample can copied to your .env file
```python
APP_HOST=127.0.0.1:3000
APP_CORS_DOMAIN=http://127.0.0.1


# LOCAL
 DATABASE_HOST=#YOUR_DB_HOST
 DATABASE_DB=#YOUR_DB_NAME
 DATABASE_USER=#YOUR_DB_USER
 DATABASE_PASSWORD=#YOUR_DB_PASSWORD
 DATABASE_PORT=#YOUR_DB_PORT
 DATABASE_SSL_MODE=disable
 DATABASE_MAX_CONNECTION=5
 DATABASE_MAX_IDLE_CONNECTION=5
 DATABASE_MAX_LIFETIME_CONNECTION=10
```
Then, go to branch development and pull request from this branch. You can run this command :
```python
git checkout -b master

git pull origin master

Enter passphrase for key '/Users/macbook/.ssh/id_rsa': 'enter your password'
```

After you open this project and pull from branch development, go to folder "server" to run this project, you can run this command :

```sh
cd server

go run main.go
```

If using file zip, you can make file .env and go to folder server and run this command
```sh
mkdir .env
```
Copy all from .env-example and filled your database

```sh
cd server

go run main.go
```


This is script to execute table in database, this project using postgresql.
```
There is in folder files - db.sql
```
