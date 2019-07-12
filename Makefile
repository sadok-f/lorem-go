PROJECT?=github.com/sadok-f/lorem-go
APP?=lorem-go
PORT?=3000

RELEASE?=0.0.2
CONTAINER_IMAGE?=docker.io/sadokf/${APP}

GOOS?=linux
GOARCH?=amd64

clean:
	rm -f ${APP}

build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build \
		-ldflags "-s -w " \
		-o ${APP}

container: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

run: container
	docker stop $(APP):$(RELEASE) || true && docker rm $(APP):$(RELEASE) || true
	docker run --name ${APP} -p ${PORT}:${PORT} --rm \
		$(CONTAINER_IMAGE):$(RELEASE)

push: container
	docker push $(CONTAINER_IMAGE):$(RELEASE)