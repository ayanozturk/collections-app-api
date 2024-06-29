FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /collections-api main.go

# Run the Go API
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /collections-api .

ENV MYSQL_DSN="app:app@tcp(db:3306)/app"

CMD ["./collections-api"]
