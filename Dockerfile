FROM golang:1.8
MAINTAINER mlabouardy <mohamed@labouardy.com>
WORKDIR /go/src/github.com/mlabouardy/swaggymnia/
COPY . .
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o swaggymnia .
ENTRYPOINT ["./swaggymnia"]
