# Use Golang base image
FROM golang:1.23.2

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to leverage caching
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application from the cmd package
RUN go build -o main ./cmd

# Command to run the application
CMD ["./main"]
