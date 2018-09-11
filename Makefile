

build: bin vendor
	go build -o bin/minasan .

bin:
	mkdir -p bin

vendor:
	dep ensure

clean:
	rm -rf bin vendor