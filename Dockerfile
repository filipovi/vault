FROM golang:1.10 AS builder
WORKDIR /go/src/github.com/filipovi/vault
COPY . /go/src/github.com/filipovi/vault
RUN  go get -d -v github.com/rs/cors && \
  go get -d -v goji.io && \
  go get -d -v gopkg.in/redis.v5 && \
  go get -d -v github.com/urfave/negroni && \
  go get -d -v github.com/filipovi/redis
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vault github.com/filipovi/vault

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/filipovi/vault .
COPY --from=builder /go/src/github.com/filipovi/vault/docker/config.json .
CMD ["./vault"]
