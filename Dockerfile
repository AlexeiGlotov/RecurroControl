FROM node:alpine3.19 AS react-build

COPY frontapp /frontapp/
WORKDIR /frontapp/
RUN npm install
RUN npm run build

FROM golang:1.21-alpine3.19 AS go-build

COPY . /recurro/
WORKDIR /recurro/

RUN go mod download
RUN go build -o ./bin/bot cmd/app/main.go

FROM nginx:alpine AS react-nginx
COPY --from=react-build /frontapp/build /usr/share/nginx/html
EXPOSE 80

FROM alpine:latest AS go-final
WORKDIR /root/
COPY --from=go-build /recurro/bin/bot .
COPY --from=go-build /recurro/configs configs/
EXPOSE 8065
CMD ["./bot"]
