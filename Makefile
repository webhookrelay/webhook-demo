TAG	= 0.0.15

build:
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo  -ldflags  -'w' -o webhook-demo .

image: build
	docker build -t karolisr/webhook-demo:$(TAG) -f Dockerfile .

push: image
	docker push	karolisr/webhook-demo:$(TAG)

quay: build
	docker build -t quay.io/rusenask/webhook-demo:$(TAG) -f Dockerfile .	