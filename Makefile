all: server

TAG = v0.1.0
PREFIX = inwinstack/ipvs-elector

.PHONY: server
server:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o server cmd/main.go

.PHONY: build_image
build_image: server
	docker build -t $(PREFIX):$(TAG) .

.PHONY: push_image
push_image:
	docker push $(PREFIX):$(TAG)

clean:
	@rm -f server