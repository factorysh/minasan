GIT_VERSION?=$(shell git describe --tags --always --abbrev=42 --dirty)

build: bin
	go build \
		-o bin/minasan \
		-ldflags "-X github.com/factorysh/minasan/version.version=$(GIT_VERSION)" \
		.

bin:
	mkdir -p bin
	chmod 777 bin

clean:
	rm -rf bin

mailhog:
	docker run --rm -p 1025:1025 -p 8025:8025 -d mailhog/mailhog

pull:
	docker pull bearstech/golang-dev
	docker pull bearstech/upx
	docker pull alpine:latest
	docker pull mailhog/mailhog

docker-build: bin vendor
	docker run --rm \
		-v ~/.cache:/.cache \
		-v `pwd`:/go/src/github.com/factorysh/minasan \
		-w /go/src/github.com/factorysh/minasan \
		-u `id -u` \
		bearstech/golang-dev \
		make build
	docker run --rm \
		-v `pwd`/bin:/upx \
		-u `id -u` \
		bearstech/upx \
		upx minasan

docker-static: bin vendor
	docker run --rm \
	-e CGO_ENABLED=0 \
	-v `pwd`:/go/src/github.com/factorysh/minasan \
	-w /go/src/github.com/factorysh/minasan \
	-u root \
	bearstech/golang-dev \
	make build

image:
	docker build -t minasan .

test:
	go test -v github.com/factorysh/minasan/minasan
