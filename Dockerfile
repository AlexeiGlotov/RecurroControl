FROM golang:1.21-alpine3.19 AS go-build

COPY . /resselgo/backapp/
WORKDIR /resselgo/backapp/

RUN go mod download
RUN go build -o ./bin/app cmd/app/main.go

FROM alpine:latest AS go-final
WORKDIR /root/
COPY --from=go-build /resselgo/backapp/bin/app .
COPY --from=go-build /resselgo/backapp/configs configs/
EXPOSE 23678
CMD ["./app"]
