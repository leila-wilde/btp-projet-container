# Dockerfile for Go-Fiber API

FROM golang:1.22-alpine AS builder

WORKDIR /build

# Copy only go.mod first to cache module download step
COPY go.mod ./
RUN go mod download

# Copy the rest of the project
COPY . .

# Ensure module files are tidy so go.sum is generated
RUN go mod tidy

# Build the Go binary from the `src` directory
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./src

# final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates curl

WORKDIR /app

COPY --from=builder /build/api .

EXPOSE 8080

CMD ["./api"]
