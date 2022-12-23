#!/bin/bash

docker run -d --name be-friends -p 4444:4444 \
    -e DB_HOST=${DB_HOST} \
    -e DB_PORT=${DB_PORT} \
    -e DB_USER=${DB_USER} \
    -e DB_PASS=${DB_PASS} \
    -e DB_NAME=${DB_NAME} \
    backend-webinar