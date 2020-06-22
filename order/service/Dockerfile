FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
RUN go build -o order/service/order-service order/service/main.go

FROM alpine
COPY --from=builder /app/order/service/order-service /app/order-service
EXPOSE 8091
ENTRYPOINT [ "/app/order-service" ]