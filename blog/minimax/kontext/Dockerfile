FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o blog .

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates sqlite

COPY --from=builder /app/blog .
COPY --from=builder /app/.env .
COPY --from=builder /app/data ./data
COPY --from=builder /app/uploads ./uploads

EXPOSE 8080

CMD ["./blog"]
