FROM golang:1.9.1 as builder
WORKDIR /usr/src/app/
COPY . .
ENV CGO_ENABLED=0 GOOS=linux
RUN go get github.com/Pepegasca/goop && make dep && make build

FROM alpine:latest
WORKDIR /usr/src/app/
COPY --from=builder /usr/src/app/ /usr/src/app/
ENV MONGODB_PORT_27017_TCP_ADDR='mongo'
EXPOSE 8000
CMD ["./bin/api"]
