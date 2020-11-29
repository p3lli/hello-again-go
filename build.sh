#!/bin/bash
IMAGE_NAME=hello-again-go
DOCKERFILE=docker-files/Dockerfile-multistage
TAG=`date +%Y%m%d%H%M%S`
BUILD_ID=intermediate_builder
LOCAL_IMAGE="${IMAGE_NAME}:${TAG}"
docker build --build-arg BUILD_ID=${BUILD_ID} -f ${DOCKERFILE} -t ${LOCAL_IMAGE} .
if test $? -eq 0
then
    printf "Build done! You can now start the container as follow:\ndocker run -p 8080:8080 ${LOCAL_IMAGE}\n"
    docker image prune --force --filter label=build=${BUILD_ID}
else
    printf "ERROR building docker image!\n"
fi
