# syntax=docker/dockerfile:1
# DOESNT WORK WITH WSL YET https://github.com/githubnemo/CompileDaemon/issues/91

FROM golang:1.21-alpine AS deps

WORKDIR /app

ADD go.mod go.sum ./
RUN go mod download &&\
    go install -mod=mod github.com/githubnemo/CompileDaemon

FROM deps AS run

ADD . ./

ENTRYPOINT CompileDaemon --log-prefix="false" --directory="/app" --build="go build -o pibot cmd/pibot/main.go" --command=./pibot
