# syntax=docker/dockerfile:1

## Developer
FROM golang:1.19-buster as development

WORKDIR /usr/src/peanut

COPY . .
RUN go install github.com/cosmtrek/air@latest
RUN go mod download

## Build
FROM golang:1.19-buster as build

WORKDIR /usr/src/peanut

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./

RUN go build -o /peanut

CMD ["/peanut"]
