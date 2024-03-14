FROM golang:1.21-alpine3.19 AS builder

COPY . /auth-bot/
WORKDIR /auth-bot/

RUN go mod download
RUN go build -o ./bin/bot cmd/app/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /auth-bot/bin/bot .
COPY --from=0 /auth-bot/notif.json .
COPY --from=0 /auth-bot/configs configs/

EXPOSE 8065

CMD ["./bot"]