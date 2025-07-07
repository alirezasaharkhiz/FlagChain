FROM golang:1.20-alpine
LABEL authors="alirezasaharkhiz"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o migrate ./cmd/server/main.go
RUN go build -o server ./cmd/server/main.go
CMD ["./server"]
ENTRYPOINT ["top", "-b"]