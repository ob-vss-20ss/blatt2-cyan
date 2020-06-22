FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
RUN go build -o client2/service/client2 client2/service/main.go

FROM alpine
COPY --from=builder /app/client2/service/client2 /app/client2
EXPOSE 8091
ENTRYPOINT [ "/app/client" ]