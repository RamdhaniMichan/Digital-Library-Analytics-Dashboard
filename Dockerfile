FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main cmd/main.go

FROM alpine:3.18

WORKDIR /app

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /app/main .
COPY --from=builder /app/migrations ./migrations
COPY .env.example .env

ENV TZ=Asia/Jakarta

EXPOSE 8081

CMD ["./main"]