APP=./build/to-do
DOCKER_TAG?=to-do


build: clean
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=amd64 go build -a -installsuffix cgo \
		-o ${APP} ./cmd/main.go

docker-build:
	docker build . -t ${DOCKER_TAG}

clean:
	@[ -f ${APP} ] && rm -f ${APP} || true