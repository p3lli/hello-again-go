#!/bin/bash
IMAGE_NAME=hello-again-go
DOCKERFILE=docker-files/Dockerfile-multistage
TAG=`date +%Y%m%d%H%M%S`
LOCAL_IMAGE="${IMAGE_NAME}:${TAG}"
docker build -f ${DOCKERFILE} -t ${LOCAL_IMAGE} .
if test $? -eq 0
then
    printf "Build done!\n"
else
    printf "ERROR building docker image!\n"
fi
