#SETUP DATABASE SCHEMA
FROM postgres:latest AS db-setup


COPY init.sql /docker-entrypoint-initdb.d/

#BUILD THE APPS
FROM golang:1.23.4-alpine AS builder

ENV CGO_ENABLED=1
ENV GOOS=linux
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o library-backend ./cmd/

#DEPLOY AND RUN THE APPPS
FROM alpine:latest

RUN apk add --no-cache iputils curl

WORKDIR /app

COPY --from=builder /app/library-backend /app/library-backend

EXPOSE 3000

CMD ["/app/library-backend"]

