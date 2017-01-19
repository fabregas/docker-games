FROM golang:1.6-alpine

ADD main.go /go/src/superservice/

RUN go install superservice

ENTRYPOINT ["superservice"]
