




# frontend
FROM node:18 as build-node
WORKDIR /build

COPY . /build
RUN \
    --mount=type=cache,target=/build/src/public/node_modules \
    make node-fetch node-build




# backend

FROM golang:latest as build-go
WORKDIR /build
RUN \
    --mount=type=cache,target=/var/cache/apt \
    --mount=type=cache,target=/var/lib/apt \
    apt-get update && apt-get install --assume-yes upx

COPY . /build
COPY --from=build-node /build/src/public/dist/ /build/src/public/dist/

RUN \
    --mount=type=cache,target=/root/.cache \
    --mount=type=cache,target=/go \
    make go-build