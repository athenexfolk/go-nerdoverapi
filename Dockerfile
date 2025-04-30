# Start with the official Golang image
FROM golang:1.24.2-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod tidy

# Copy the entire Go project to the container
COPY . .

# Build the Go application
RUN go build -o main .

# Start with a smaller base image for the final image
FROM alpine:latest

# Install required libraries for running Go apps (Gin uses this)
RUN apk --no-cache add ca-certificates

# Set the working directory in the final image
WORKDIR /root/

# Copy the compiled binary from the builder image to the final image
COPY --from=builder /app/main .

# Expose the port your application will run on (usually 8080 for Gin)
EXPOSE 8080

# Command to run the app
CMD ["./main"]
