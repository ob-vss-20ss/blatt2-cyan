FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
RUN go build -o client4/service/client4 client4/service/main.go

FROM alpine
COPY --from=builder /app/client4/service/client4 /app/client4
EXPOSE 8091
ENTRYPOINT [ "/app/client" ]