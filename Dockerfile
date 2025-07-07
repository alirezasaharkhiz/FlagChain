FROM golang:1.24-alpine

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata netcat-openbsd

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY .env.example .env

COPY wait-for-db.sh /wait-for-db.sh
RUN chmod +x /wait-for-db.sh

RUN go build -o server ./cmd/server/main.go

CMD ["/wait-for-db.sh", "./server"]
