# Use an official Golang image
FROM golang:1.24.2

# Set working directory inside container
WORKDIR /app

# Copy source code to container
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./main"]