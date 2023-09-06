# Builder
FROM golang:1.20-alpine as builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy
COPY . .

RUN go build -o main ./cmd/server

# Run stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/main .
COPY credentials credentials

# Expose port 8080
EXPOSE 8080
# Set the entrypoint to run the Golang application
CMD ["/app/main"]