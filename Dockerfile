FROM golang:1.23-alpine

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY . .

RUN go build -o main ./cmd/main.go

CMD sleep 5 && ./main