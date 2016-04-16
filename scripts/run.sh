#!/bin/bash
set -x #echo on

docker run \
  -e AWS_PRIVATE_ACCESS_KEY_ID=$AWS_PRIVATE_ACCESS_KEY_ID \
  -e AWS_PRIVATE_SECRET_ACCESS_KEY=$AWS_PRIVATE_SECRET_ACCESS_KEY \
  avatarme:latest
