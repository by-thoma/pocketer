FROM golang:1.19-alpine3.16 AS builder

COPY . /github.com/by-thoma/pocketer/
WORKDIR /github.com/by-thoma/pocketer/

RUN go mod download
RUN go build -o ./.bin/bot ./cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/by-thoma/pocketer/.bin/bot .
COPY --from=0 /github.com/by-thoma/pocketer/configs configs/

EXPOSE 80

CMD ["./bot"]
