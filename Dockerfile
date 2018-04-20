FROM golang:1.9

RUN apt-get update
RUN apt-get install -y libmagickwand-dev
RUN go get gopkg.in/gographics/imagick.v2/imagick

WORKDIR /go/src/ic-app
COPY . .

RUN go build -o /go/bin/ic main.go

EXPOSE 8081

CMD ["/go/bin/ic"]