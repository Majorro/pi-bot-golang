# syntax=docker/dockerfile:1

FROM golang:1.21-alpine AS deps

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

FROM deps AS build
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /pibot cmd/pibot/main.go

FROM build AS test
RUN go test -v ./...

FROM scratch AS release

RUN apk add --update --no-cache ca-certificates git

WORKDIR /
COPY --from=build /pibot /pibot

ENTRYPOINT ["/pibot"]
