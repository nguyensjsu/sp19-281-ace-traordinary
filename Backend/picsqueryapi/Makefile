all: clean

clean:
	find . -name 'picsqueryapi' -type f -exec rm -f {} \;
	go clean

go-get:
	rm -rf src/github.com
	go get -v github.com/codegangsta/negroni
	go get -v github.com/gorilla/mux
	go get -v github.com/unrolled/render
	go get -v github.com/satori/go.uuid
	go get -v github.com/streadway/amqp
	go get -v database/sql
	go get -v github.com/go-sql-driver/mysql
	go get -v github.com/aws/aws-sdk-go
	go get -v github.com/golang/glog
	go get -v github.com/gorilla/handlers

run:
	go run src/app/$(app).go

main:
	go run src/app/main.go

format:
	go fmt picscmdapi

install:
	go install picsqueryapi

build:
	go build github.com/sp19-281-ace-traordinary/Backend/picsqueryapi/src/picsqueryapi

start:
	./picsqueryapi

docker-build:
	docker build -t 1ra4vi3/picsqueryapi .
	docker images

docker-network:
	docker network ls

docker-network-inspect:
	docker network inspect host

network-create:
	docker network create picsqueryapi

network-prune:
	docker network prune

docker-ip:
	docker-machine ip

docker-ps:
	 docker ps --all --format "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Status}}\t"

docker-ps-ports:
	 docker ps --all --format "table {{.Names}}\t{{.Ports}}\t"
