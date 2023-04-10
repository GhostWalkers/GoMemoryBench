# syntax=docker/dockerfile:1
FROM golang:1.20

WORKDIR /app
COPY ./go.mod ./
RUN go mod download
