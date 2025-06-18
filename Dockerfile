FROM golang:1.22-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# ✅ Copy file build.yaml (đảm bảo nó nằm trong thư mục gốc hoặc điều chỉnh path phù hợp)
COPY build.yaml ./config/build.yaml

RUN go build -o main .

EXPOSE 8080
EXPOSE 8082

CMD ["./main"]
