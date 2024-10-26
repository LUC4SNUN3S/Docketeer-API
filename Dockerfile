FROM golang:1.22.1 as builder

WORKDIR /app

COPY . .

RUN go build -o tarelai-webhook-receive /app/cmd/tarelai-webhook-receive/

FROM alpine

WORKDIR /app

COPY --from=builder /app/tarelai-webhook-receive /app/tarelai-webhook-receive

EXPOSE 8080

# Comando para rodar a aplicação

ENTRYPOINT ["/docketeer-api"]