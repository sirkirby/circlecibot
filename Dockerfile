FROM golang:1.10.1

RUN apt-get update

ADD . /go/src/github.com/sirkirby/circlebot

WORKDIR /go/src/github.com/sirkirby/circlebot 

RUN go get github.com/golang/dep/cmd/dep && dep ensure

RUN cd cmd/slash && go build

CMD ["cmd/slash/slash"]