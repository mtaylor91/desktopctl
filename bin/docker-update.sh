#!/bin/sh
IMAGE=images.home.mtaylor.io/desktopctl
docker build -t $IMAGE .
docker push $IMAGE
