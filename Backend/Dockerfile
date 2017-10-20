FROM golang
EXPOSE 8000
RUN mkdir /app
WORKDIR /app
ADD bin bin
ADD config config
ADD src src
ADD api_app.go api_app.go
ADD Goopfile Goopfile
ADD Makefile Makefile
RUN go get github.com/Pepegasca/goop
RUN make dep
RUN make build
ENTRYPOINT ["./bin/api"] 