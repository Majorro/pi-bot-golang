# syntax=docker/dockerfile:1

FROM golang:1.21-alpine AS deps

RUN apk --update add ca-certificates git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

FROM deps AS build
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /pibot cmd/pibot/main.go

FROM build AS test
RUN go test -v ./...

FROM scratch AS release

WORKDIR /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /pibot /pibot

ENTRYPOINT ["/pibot"]
