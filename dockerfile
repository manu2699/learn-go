# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the entire application source code
COPY . .

# Download and install the Go dependencies
RUN go get

# Build the Go application
RUN go build -o server .

# Docker will execute the "/app/server" binary when container starts.
ENTRYPOINT ["/app/server"]
