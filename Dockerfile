FROM golang:1.21-alpine3.19 AS go-build

COPY . /recurro/backapp/
WORKDIR /recurro/backapp/

RUN go mod download
RUN go build -o ./bin/app cmd/app/main.go

FROM alpine:latest AS go-final
WORKDIR /root/
COPY --from=go-build /recurro/backapp/bin/app .
COPY --from=go-build /recurro/backapp/configs configs/
EXPOSE 23678
CMD ["./app"]
