#!/usr/bin/env bash

if [ -z $1 ]; then
  echo "ERROR: No docker hub user provided."
  echo " Syntax: "$0" username vX.YY"
  exit 2
fi

if [ -z $2 ]; then
  echo "ERROR: No version provided."
  echo " Syntax: "$0 $1" vX.YY"
  exit 2
fi

APP="cns_stepsdatafeed"
VERSION=$2
DOCKERUSER=$1

ID=$(docker build -q -t $APP:$VERSION .)
echo "- docker image built as $ID"
docker tag $ID $DOCKERUSER/$APP:$VERSION
echo "- docker image $ID tagged as $DOCKERUSER/$APP:$VERSION"
docker push $DOCKERUSER/$APP:$VERSION
echo "- docker image pushed to https://hub.docker.com/repository/docker/$DOCKERUSER/$APP , version $VERSION"
