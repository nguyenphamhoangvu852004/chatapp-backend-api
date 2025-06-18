FROM golang:1.24-alpine

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

RUN go build -o bin/main ./cmd/server

CMD ["./bin/main"]
