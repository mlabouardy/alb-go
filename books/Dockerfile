FROM golang:1.8.0
LABEL maintainer mlabouardy

ENV DIR /go/src/github.com/mlabouardy/app/

COPY . $DIR

WORKDIR $DIR

EXPOSE 5000

CMD go run app.go
