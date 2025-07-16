FROM golang:1.23.0-alpine
LABEL authors="Log1c0"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o cmd/paymentapi/paymentapi cmd/paymentapi/main.go

EXPOSE 8080

CMD ["./cmd/paymentapi/paymentapi"]
