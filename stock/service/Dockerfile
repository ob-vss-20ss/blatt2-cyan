FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
RUN go build -o stock/service/stock-service stock/service/main.go

FROM alpine
COPY --from=builder /app/stock/service/stock-service /app/stock-service
EXPOSE 8091
ENTRYPOINT [ "/app/stock-service" ]