# step 1: build executable binary
FROM golang:1.19.1-alpine3.16 AS builder
RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /agmc-day-6

# step 2: build a small image
FROM alpine:3.16.0
WORKDIR /app
COPY --from=builder agmc-day-6 .
EXPOSE $PORT
CMD ["./agmc-day-6", "-m=migrate"]