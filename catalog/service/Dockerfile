FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
RUN go build -o catalog/service/catalog-service catalog/service/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/catalog/service/catalog-service /app/catalog-service
COPY data/catalog.json /app/data/catalog.json
EXPOSE 8091
ENTRYPOINT [ "/app/catalog-service" ]