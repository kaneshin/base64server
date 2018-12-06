FROM golang:1.11.2

RUN mkdir -p /go/src/github.com/kaneshin/base64server
ADD . /go/src/github.com/kaneshin/base64server/

RUN go build -o main github.com/kaneshin/base64server

EXPOSE 8080
CMD ["./main"]
