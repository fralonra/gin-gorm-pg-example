FROM golang:1.9
WORKDIR /go/src/gin-gorm-pg-example
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
EXPOSE 8888
ENTRYPOINT ["gin-gorm-pg-example"]