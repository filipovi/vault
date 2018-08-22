FROM golang:1.10 AS builder
WORKDIR /go/src/github.com/filipovi/vault
COPY . /go/src/github.com/filipovi/vault
RUN go get github.com/micro/go-micro
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vault /go/src/github.com/filipovi/vault/server/server.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/filipovi/vault .
CMD ["./vault"]

