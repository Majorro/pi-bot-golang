# syntax=docker/dockerfile:1

FROM golang:1.21-alpine AS build

WORKDIR /app

COPY . ./
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /pibot cmd/pibot/main.go

FROM build AS test
RUN go test -v ./...

# not scratch because issues with certificates
FROM alpine:latest AS release

# TODO: remove?
RUN apk add libc6-compat

WORKDIR /
COPY --from=build /pibot /pibot

ENTRYPOINT ["/pibot"]
