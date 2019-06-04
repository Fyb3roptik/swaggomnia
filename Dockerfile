FROM golang:1.11
MAINTAINER Fyb3roptik <nwallace@fyberstudios.com>
WORKDIR /go/src/github.com/fyb3roptik/swaggomnia/
COPY . .
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o swaggomnia .
ENTRYPOINT ["./swaggomnia"]
