FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
RUN go build -o customer/service/customer-service customer/service/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/customer/service/customer-service /app/customer-service
COPY data/customers.json /app/data/customers.json
EXPOSE 8091
ENTRYPOINT [ "/app/customer-service" ]