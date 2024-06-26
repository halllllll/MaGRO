include .env

build:
	time DOCKER_BUILDKIT=1 docker build -f ./docker/Dockerfile.prod \
	--platform ${PROD_PLATFORM} \
	--progress=plain \
	--build-arg FRONT_IMAGE=${BUN_IMAGE} \
	--build-arg SERVER_1STSTAGE_IMAGE=${GO_DEV_IMAGE} \
  --build-arg SERVER_PROD_IMAGE=${GO_PROD_IMAGE} \
	--build-arg GO_OS=${PROD_OS} \
	--build-arg GO_ARCH=${PROD_ARCH} \
	--build-arg GO_PROD_PLATFORM=${PROD_PLATFORM} \
	--build-arg EXPOSE_PORT=${CONTAINER_PORT} \
	--no-cache \
	--force-rm \
	-t ${IMAGE_NAME}:latest . \
	&& docker image prune -f
ß
save:
	docker save ${IMAGE_NAME}:latest -o app.tar

# ローカル確認用(ポートは好きに変えてね)
run:
	docker container run --rm -p 5522:${CONTAINER_PORT} -e GO_APP_PORT=${CONTAINER_PORT} ${IMAGE_NAME}

## 実行する場合サンプル。公開するときは`-p　8080:10201`なんかも必要かも
## docker run --rm --name test --expose 10201 -e GO_APP_PORT=10201 MaGRO
