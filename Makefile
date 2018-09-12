

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