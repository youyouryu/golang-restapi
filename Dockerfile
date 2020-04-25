FROM golang

WORKDIR /go/src/github.com/restapi
ADD . .
RUN go install
ENTRYPOINT /go/bin/restapi
EXPOSE 8080

