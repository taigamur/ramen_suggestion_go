FROM golang:1.17.6-alpine
RUN apk update && apk add git
WORKDIR /go/src

CMD ["go", "run", "main.go"]

