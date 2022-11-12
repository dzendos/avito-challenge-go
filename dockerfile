FROM golang:latest AS builder
WORKDIR /build
COPY . .
RUN go build -o /build/app github.com/dzendos/avito-challenge/cmd/service