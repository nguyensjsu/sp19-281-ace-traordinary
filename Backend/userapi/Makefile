
all: clean

clean:
	find . -name 'userapi' -type f -exec rm -f {} \;
	go clean

run:
	go run $(app).go

main:
	go run main.go

format:
	go fmt userapi

install:
	go install userapi

build:
	go build userapi
start:
		./userapi

docker-build:
	docker build -t userapi .
	docker images

network-create:
	docker network create --driver bridge userapi

docker-run:
	docker run --network userapi \
				--name userapi -p 3000:3000 -td userapi
	docker ps

docker-network:
	docker network ls

docker-network-prune:
	docker network prune

docker-network-inspect:
	docker network inspect host

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

test-ping:
		 	curl localhost:3000/ping
