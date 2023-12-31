# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Go application
RUN go build -o app ./cmd/server

# Expose the port your application listens on
EXPOSE 4000

# Command to run your application
CMD ["./app"]
