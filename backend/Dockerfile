FROM golang:1.23.3 AS build-stage

WORKDIR /app
COPY . .

ARG ENV_FILE=.env
COPY ${ENV_FILE} .env

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server cmd/server/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=build-stage /app/server .
COPY --from=build-stage /app/.env .env

RUN chmod +x /app/server

EXPOSE 8085

CMD ["/app/server"]