FROM golang:1.12.4-alpine3.9

ADD . /go/src/github.com/probably-not/go-application-example

WORKDIR /go/src/github.com/probably-not/go-application-example

RUN go build -tags=jsoniter ./cmd/gin-api-example

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
RUN mkdir configs
COPY --from=0 /go/src/github.com/probably-not/go-application-example/gin-api-example ./gin-api-example
COPY --from=0 /go/src/github.com/probably-not/go-application-example/configs ./configs

EXPOSE 8080

CMD ["./api"]