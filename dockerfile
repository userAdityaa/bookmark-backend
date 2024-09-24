# Start from Golang base image
FROM golang:1.18-alpine

# Set environment variables
ENV GO111MODULE=on

# Create an app directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go app
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Start the Go app
CMD ["./main"]
