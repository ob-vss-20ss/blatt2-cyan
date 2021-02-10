FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
RUN go build -o payment/service/payment-service payment/service/main.go

FROM alpine
COPY --from=builder /app/payment/service/payment-service /app/payment-service
EXPOSE 8091
ENTRYPOINT [ "/app/payment-service" ]