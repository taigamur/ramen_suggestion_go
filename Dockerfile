FROM golang:1.17.6-alpine
RUN apk update && apk add git

WORKDIR /go/src
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./

RUN go build -o /main

CMD ["/main"]