proto:
	bash ./gen-proto.sh

build-organizations-go:
	cd services/organizations && go get ./... && go build -o run && cd ../..

build-scenes-go:
	cd services/scenes && go get ./... && go build -o run && cd ../..

build-operators-go:
	cd services/operators && go get ./... && go build -o run && cd ../..

build-parameters-go:
	cd services/parameters && go get ./... && go build -o run && cd ../..

build-api-go:
	cd services/api && go get ./... && go build -o run && cd ../..

build-compute-go:
	cd services/compute && go get ./... && go build -o run && cd ../..

build-go: build-api-go build-compute-go


build-api:
	docker-compose build server

build-compute:
	docker-compose build compute

build: proto vendor build-server build-compute

up: build
	docker-compose up server compute

test:
	bash ./run-tests.integration.sh

migrate-cockroachdb:
	migrate -database
