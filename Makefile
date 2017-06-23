API_LIB_FILES = $(shell find applications/api/lib -type f -name "*.go")
API_SERVICES_FILES = $(shell find applications/api/services/web -type f -name "*.go")
API_TASKS_FILES = $(shell find applications/api/tasks -type f -name "*.go")
COMPUTE_SERVICES_FILES = $(shell find applications/compute/services/web -type f -name "*.go")
CORE_FILES = $(shell find core -type f -name "*.go")
PROTO_FILES = $(shell find protos -type f -name "*.proto")

# First target is default
build-go: fmt build-api-go build-compute-go

clean:
	touch applications/api/services/web/run && rm applications/api/services/web/run
	touch applications/compute/services/web/run && rm applications/compute/services/web/run

configure:
	go get -u github.com/jteeuwen/go-bindata/...

fmt:
	go fmt ./applications/...
	go fmt ./core/...

#############
# Go binaries
#############

applications/api/tasks/migrate/bindata.go: $(shell find applications/api/tasks/migrate/sql -name "*.sql")
	# Using -prefix "sql/" because go-bindata postgres filename parsing can't handle path prefix
	cd applications/api/tasks/migrate && go-bindata -pkg migrate -prefix "sql/" sql
bindata: applications/api/tasks/migrate/bindata.go

applications/api/services/web/run: $(API_SERVICES_FILES) $(API_LIB_FILES) $(CORE_FILES)
	cd applications/api/services/web && go get ./... && GOOS=linux GOARCH=amd64 go build -o run
applications/api/tasks/run: $(API_TASKS_FILES) $(API_LIB_FILES)
	cd applications/api/tasks && go get ./... && GOOS=linux GOARCH=amd64 go build -o run
build-api-go: fmt applications/api/services/web/run applications/api/tasks/run

applications/compute/services/web/run: $(COMPUTE_SERVICES_FILES) $(CORE_FILES)
	cd applications/compute/services/web && go get ./... && GOOS=linux GOARCH=amd64 go build -o run
build-compute-go: fmt applications/compute/services/web/run

###############
# Docker images
###############

build-api: all-protos build-api-go
	docker-compose build api-services

build-compute: all-protos build-compute-go
	docker-compose build compute-services

build: build-api build-compute

.PHONY: bootstrap-postgres
bootstrap-postgres: build-api
	./scripts/bootstrap-postgres.sh

#####################
# Generate proto code
#####################

.PHONY: protoman
protoman:
	docker build -t stupschwartz/protoman -f protoman/Dockerfile protoman

all-protos: $(PROTO_FILES)
	./scripts/generate-protos.sh

proto-gen/services/%.pb: protos/%.proto
	./scripts/generate-protos.sh $* service

proto-gen/go/%.pb.go: protos/%.proto
	./scripts/generate-protos.sh $* go

proto-gen/js/%_pb.js proto-gen/js/%_grpc_pb.js: protos/%.proto
	./scripts/generate-protos.sh $* js

protonames = $(shell find protos -type f -name "*.proto" | xargs -n1 basename | awk '{split($$0,a,"."); print a[1]}')

protos: $(foreach protoname,$(protonames),$(subst NAME,$(protoname),proto-gen/services/NAME/NAME.pb proto-gen/go/NAME/NAME.pb.go proto-gen/js/NAME/NAME_pb.js))

################
# Run containers
################

up: build
	docker-compose up server compute

.PHONY: test
test:
	./run-tests.integration.sh

migrate-cockroachdb:
	migrate -database
