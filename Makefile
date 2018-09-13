

build: bin vendor
	go build -o bin/minasan .

bin:
	mkdir -p bin
	chmod 777 bin

vendor:
	dep ensure

clean:
	rm -rf bin vendor

mailhog:
	docker run --rm -p 1025:1025 -p 8025:8025 -d mailhog/mailhog

docker-pull:
	docker pull bearstech/golang-dep
	docker pull bearstech/upx

docker-build: bin
	docker run --rm \
	-v `pwd`:/go/src/gitlab.bearstech.com/factory/minasan \
	-w /go/src/gitlab.bearstech.com/factory/minasan \
	-u root \
	bearstech/golang-dep \
	make build
	docker run --rm \
	-v `pwd`/bin:/upx \
	bearstech/upx \
	upx minasan

