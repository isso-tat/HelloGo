# Production environment (alias:base)
FROM golang:1.12-alpine as base
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
WORKDIR /home/HelloGo
RUN ls
RUN ls

# Dev environment (alias:dev)
# FROM base as dev
