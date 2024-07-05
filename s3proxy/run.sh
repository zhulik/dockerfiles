#!/bin/sh

set -eu

docker build -t ghcr.io/zhulik/s3_proxy .
docker run -p 9292:9292 \
       -e ENDPOINT=$ENDPOINT \
       -e AWS_REGION=$AWS_REGION \
       -e BUCKET=$BUCKET \
       -e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID \
       -e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY \
       -e FORCE_PATH_STYLE=1 \
       -it ghcr.io/zhulik/s3_proxy
