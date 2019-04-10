
all: clean

clean:
	find . -name 'userapi' -type f -exec rm -f {} \;
	go clean

run:
	go run src/app/$(app).go

main:
	go run src/app/main.go

format:
	go fmt goapi

install:
	go install goapi

build:
	go build goapi

start:
	./goapi

test-ping:
	curl localhost:3000/ping

docker-build:
	docker build -t userapi .
	docker images

docker-run:
	docker run --name userapi -td -p 3000:3000 userapi
	docker ps

docker-network:
	docker network ls

docker-network-inspect:
	docker network inspect host

docker-shell:
	docker exec -it userapi bash

docker-clean:
	docker stop userapi
	docker rm userapi
	docker rmi userapi

docker-ip:
	docker-machine ip

docker-ps:
	 docker ps --all --format "table {{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Status}}\t"

docker-ps-ports:
	 docker ps --all --format "table {{.Names}}\t{{.Ports}}\t"
