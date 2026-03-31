# Build stage
FROM golang:1.21-alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=1 go build -o blog-backend .

# Runtime stage
FROM alpine:3.19

RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/blog-backend .

RUN mkdir -p /app/uploads

EXPOSE 8080

CMD ["./blog-backend"]
