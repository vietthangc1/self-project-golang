FROM golang:1.20-alpine
WORKDIR /app

# Copy the current directory contents into the container at /app

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy
COPY . .

# Build the Golang application
RUN go build ./cmd/server

# Expose port 8080
EXPOSE 3000

# Set the entrypoint to run the Golang application
CMD ["go", "run", "./cmd/server"]
