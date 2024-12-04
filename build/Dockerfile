FROM golang:1.23 as builder

WORKDIR /github.com/babyhando/order-service

RUN CGO_ENABLED=0 go build -o ./order-api cmd/main.go

FROM alpine as deploy

RUN apk add --no-cache tzdata
ENV TZ=Asia/Tehran

WORKDIR /github.com/babyhando/order-service

COPY --from=builder /github.com/babyhando/order-service/order-api ./order-api

CMD [ "./order-api" , "--config" , "/etc/config.json"]