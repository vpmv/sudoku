FROM golang:1.15-alpine
RUN apk update \
    && apk --no-cache upgrade \
    && apk --no-cache add \
        git \
;

RUN mkdir -p /app/bin/mac

COPY src /app/src
WORKDIR /app/src
RUN go mod download

# build different solvers for linux & mac
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /app/bin/bruteforce
RUN CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o /app/bin/mac/bruteforce