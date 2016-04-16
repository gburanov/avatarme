#!/bin/bash
set -x #echo on

GOOS=linux go build
docker build -t avatarme .
