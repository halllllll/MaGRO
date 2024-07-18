include .env

.PHONY: upserter

# いったんローカルテストで外す用
#  --platform ${PROD_PLATFORM} \

build:
	time DOCKER_BUILDKIT=1 docker build -f ./docker/Dockerfile.prod \
	--platform ${PROD_PLATFORM} \
	--progress=plain \
	--build-arg FRONT_IMAGE=${FRONT_PROD_IMAGE} \
	--build-arg SERVER_1STSTAGE_IMAGE=${GO_DEV_IMAGE} \
  --build-arg SERVER_IMAGE=${GO_PROD_IMAGE} \
	--build-arg OS=${PROD_OS} \
	--build-arg ARCH=${PROD_ARCH} \
	--build-arg PLATFORM=${PROD_PLATFORM} \
	--build-arg ENTRA_CLIENT_ID=${ENTRA_CLIENT_PROD_ID} \
	--build-arg ENTRA_TENANT_ID=${ENTRA_TENANT_PROD_ID} \
	--build-arg ENTRA_REDIRECT_URI=${ENTRA_REDIRECT_PROD_URI} \
	--build-arg CLIENT_SECRET=${ENTRA_CLIENT_PROD_SECRET} \
	--build-arg URI=${PROD_URI} \
	--build-arg GO_APP_PORT=${GO_APP_PROD_PORT} \
	--no-cache \
	--force-rm \
	-t ${IMAGE_NAME}:latest . \
	&& docker image prune -f


# upserter実行 db名, network名, service名は合わせて変更
upserter:
	docker container run --rm -t -v ./upserter-data:/files -e DBPORT=${DB_PORT} -e  DB_USER=${DB_USER} -e DB_PASSWORD=${DB_PASSWORD} -e DB_HOSTNAME=db -e DB_NAME=${DB_NAME} -e SERVICE=lgate --network magro_default   magro-upserter

# ローカルテスト用
# 	--build-arg IMAGE=golang:1.22.4-bullseye \
#		--build-arg IMAGE=${GO_PROD_IMAGE} \

build-upseter:
	time DOCKER_BUILDKIT=1 docker build -f ./docker/Dockerfile.upserter \
	--platform ${PROD_PLATFORM} \
	--progress=plain \
	--build-arg PLATFORM=${PROD_PLATFORM} \
	--build-arg BUILDER_IMAGE=${GO_DEV_IMAGE} \
	--build-arg PROD_IMAGE=${GO_PROD_IMAGE} \
	--build-arg OS=${PROD_OS} \
	--build-arg ARCH=${PROD_ARCH} \
	--no-cache \
	--force-rm \
	-t magro-upserter:latest . \
	&& docker image prune -f

save:
	docker save ${IMAGE_NAME}:latest -o app.tar

# ローカル確認用(ポートは好きに変えてね)
run:
	docker container run --rm -p 5522:${CONTAINER_PORT} -e GO_APP_PORT=${CONTAINER_PORT} ${IMAGE_NAME}

## 実行する場合サンプル。公開するときは`-p　8080:10201`なんかも必要かも
## docker run --rm --name test --expose 10201 -e GO_APP_PORT=10201 MaGRO
