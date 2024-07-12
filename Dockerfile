FROM golang:1.23rc1-alpine

WORKDIR /app

COPY go.mod .

RUN go mod tidy

COPY . .

RUN go build -o main ./cmd/api/main.go

EXPOSE 8080

CMD ["./main"]
