FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/api/main.go

FROM golang:1.21-alpine as dev
WORKDIR /app
COPY --from=builder /app/main .
RUN apk add --no-cache git
RUN go install github.com/cosmtrek/air@latest
EXPOSE 5000
CMD ["air"]

FROM scratch as prod
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 5000
CMD ["./main"]