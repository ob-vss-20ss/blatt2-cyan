FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
RUN go build -o stock/service/stock-service stock/service/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/stock/service/stock-service /app/stock-service
COPY data/stock.json /app/data/stock.json
EXPOSE 8091
ENTRYPOINT [ "/app/stock-service" ]