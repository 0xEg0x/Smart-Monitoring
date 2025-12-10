
FROM golang:1.22-alpine AS builder

RUN apk add --no-cache git


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app .


FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/app .

ENV N8N_WEBHOOK="http://n8n:5678/webhook/monitor"

# Expor porta  servidor Go
EXPOSE 8080

# Rodar app
CMD ["./app"]
