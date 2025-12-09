FROM golang:1.20-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app main.go

FROM alpine:3.18

WORKDIR /app
COPY --from=build /app/app .

CMD ["./app"]
