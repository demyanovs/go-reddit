# go-reddit

Download
https://github.com/golang-migrate/migrate/releases

For linux:
tar -xf migrate.linux-amd64.tar.gz
sudo mv migrate.linux-amd64 /usr/local/bin/migrate

To start database
make postgres

Open
http://localhost:8080/

go get github.com/jmoiron/sqlx

go get github.com/lib/pq

## chi router
go get -u github.com/go-chi/chi

## Run server
go run cmd/goreddit/main.go 