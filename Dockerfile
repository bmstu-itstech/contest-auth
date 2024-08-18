FROM golang:1.22.5-alpine AS builder

COPY app /github.com/bmstu-itstech/contest-auth/source
WORKDIR /github.com/bmstu-itstech/contest-auth/source

RUN go mod download
RUN go build -o ./bin/contest-auth cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/bmstu-itstech/contest-auth/source/bin/contest-auth .

ADD config/config.yaml /config.yaml

CMD ["./contest-auth"]