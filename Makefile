TAG	= latest

build:
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo  -ldflags  -'w' -o webhook-demo .

image: build
	docker build -t karolisr/webhook-demo:$(TAG) -f Dockerfile .

push: image
	docker push	karolisr/webhook-demo:$(TAG)

private: build
	docker build -t karolisr/demo-webhook:$(TAG) -f Dockerfile .
	docker push karolisr/demo-webhook:$(TAG)

ecr: build
	docker build -t 528670773427.dkr.ecr.us-east-2.amazonaws.com/webhook-demo:$(TAG) -f Dockerfile .
	docker push 528670773427.dkr.ecr.us-east-2.amazonaws.com/webhook-demo:$(TAG)