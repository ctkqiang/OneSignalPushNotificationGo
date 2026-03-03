FROM golang:1.24.5-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pushservice .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/pushservice .
COPY --from=builder /app/internal/config ./internal/config
COPY --from=builder /app/docs ./docs
EXPOSE 8080
CMD ["./pushservice", "--release"]