FROM obraun/golang-micro-ci as builder
COPY . /app
WORKDIR /app
RUN go build -o client3/service/client3 client3/service/main.go

FROM alpine
COPY --from=builder /app/client3/service/client3 /app/client3
EXPOSE 8091
ENTRYPOINT [ "/app/client3" ]