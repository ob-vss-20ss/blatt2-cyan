FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
RUN go build -o client1/service/client1 client1/service/main.go

FROM alpine
COPY --from=builder /app/client1/service/client1 /app/client1
EXPOSE 8091
ENTRYPOINT [ "/app/client" ]