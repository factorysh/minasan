

build: bin vendor
	go build -o bin/minasan .

bin:
	mkdir -p bin

vendor:
	dep ensure

clean:
	rm -rf bin vendor

mailhog:
	docker run --rm -p 1025:1025 -p 8025:8025 -d mailhog/mailhog

docker-pull:
	docker pull bearstech/golang-dev
	docker pull bearstech/upx

docker-build:
	docker run --rm \
	-v `pwd`:/go/src/gitlab.bearstech.com/factory/minasan \
	-w /go/src/gitlab.bearstech.com/factory/minasan \
	bearstech/golang-dev \
	make build
	docker run --rm \
	-v `pwd`/bin:/upx \
	bearstech/upx \
	upx minasan

