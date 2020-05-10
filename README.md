# go-reddit

##Database migrations

Download
``
https://github.com/golang-migrate/migrate/releases
``

For linux:

``
tar -xf migrate.linux-amd64.tar.g
``

``
sudo mv migrate.linux-amd64 /usr/local/bin/migrate
``

##Dependencies

go get github.com/jmoiron/sqlx

go get github.com/lib/pq

go get -u github.com/go-chi/chi

##Start database

``
make postgres
``

Open
http://localhost:8080/

## Run server
go run cmd/goreddit/main.go 