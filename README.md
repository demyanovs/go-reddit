# go-reddit

A simple reddit clone/style in Go. 

Based on gowebexamples.com - <a href="https://course.gowebexamples.com/course/creating-a-reddit-clone-in-go" target="_blank">Creating a reddit clone in Go</a>. 

## Database migrations

Download
````
https://github.com/golang-migrate/migrate/releases
````

For linux:

````
tar -xf migrate.linux-amd64.tar.g
````

````
sudo mv migrate.linux-amd64 /usr/local/bin/migrate
````

## Dependencies

go get github.com/jmoiron/sqlx

go get github.com/lib/pq

go get -u github.com/go-chi/chi

cd /tmp && 
go get github.com/cespare/reflex && 
go install github.com/cespare/reflex

## Start database

````
make postgres
````

Open
http://localhost:8080/

## Run server
````
go run cmd/goreddit/main.go
````  

Autoreloading with reflex
````
reflex -s go run cmd/goreddit/main.go
````

````
http://localhost:3000/
````