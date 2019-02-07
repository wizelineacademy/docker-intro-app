FROM golang:1.9.1 as docker-academy-backend-builder
RUN mkdir /app
WORKDIR /app
COPY . .
ENV CGO_ENABLED=0 GOOS=linux
RUN go get github.com/Pepegasca/goop && make dep && make build

FROM alpine:latest
RUN apk -U add curl && \
    rm -rf /var/cache/apk/*
COPY --from=docker-academy-backend-builder app app
WORKDIR /app
EXPOSE 8000
CMD ["./bin/api"]
