FROM golang:1.23 as builder

WORKDIR /order-service

RUN CGO_ENABLED=0 go build -o ./order-api cmd/main.go

FROM alpine as deploy

RUN apk add --no-cache tzdata
ENV TZ=Asia/Tehran

WORKDIR /order-service

COPY --from=builder /order-service/order-api ./order-api

CMD [ "./order-api" , "--config" , "/etc/config.json"]